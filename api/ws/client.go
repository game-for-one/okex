package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/game-for-one/okex"
	"github.com/game-for-one/okex/events"
	"github.com/gorilla/websocket"
)

// ClientWs is the websocket api client
//
// https://www.okex.com/docs-v5/en/#websocket-api
type ClientWs struct {
	Cancel              context.CancelFunc
	DoneChan            chan interface{}
	StructuredEventChan chan interface{}
	RawEventChan        chan *events.Basic
	ErrChan             chan *events.Error
	WSErrChan           chan error
	SubscribeChan       chan *events.Subscribe
	UnsubscribeCh       chan *events.Unsubscribe
	LoginChan           chan *events.Login
	SuccessChan         chan *events.Success
	PrivateSendChan     chan []byte
	PublicSendChan      chan []byte
	url                 map[bool]okex.BaseURL
	PrivateConn         *websocket.Conn
	PrivateConnMutex    *sync.RWMutex
	PublicConn          *websocket.Conn
	PublicConnMutex     *sync.RWMutex
	apiKey              string
	secretKey           []byte
	passphrase          string

	PrivateWriteMutex *sync.Mutex
	PublicWriteMutex  *sync.Mutex

	lastTransmitMutex *sync.RWMutex
	lastTransmit      map[bool]*time.Time
	// mu                  map[bool]*sync.RWMutex
	AuthRequested *time.Time
	Authorized    bool
	Private       *Private
	Public        *Public
	Trade         *Trade
	ctx           context.Context
}

const (
	redialTick = 2 * time.Second
	writeWait  = 3 * time.Second
	pongWait   = 30 * time.Second
	PingPeriod = (pongWait * 8) / 10
)

// NewClient returns a pointer to a fresh ClientWs
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, url map[bool]okex.BaseURL) *ClientWs {
	ctx, cancel := context.WithCancel(ctx)
	c := &ClientWs{
		apiKey:              apiKey,
		secretKey:           []byte(secretKey),
		passphrase:          passphrase,
		ctx:                 ctx,
		Cancel:              cancel,
		url:                 url,
		PrivateSendChan:     make(chan []byte, 3),
		PublicSendChan:      make(chan []byte, 3),
		PublicConnMutex:     &sync.RWMutex{},
		PrivateConnMutex:    &sync.RWMutex{},
		PrivateWriteMutex:   &sync.Mutex{},
		PublicWriteMutex:    &sync.Mutex{},
		DoneChan:            make(chan interface{}),
		StructuredEventChan: make(chan interface{}),
		RawEventChan:        make(chan *events.Basic),
		lastTransmit:        make(map[bool]*time.Time),
		lastTransmitMutex:   &sync.RWMutex{},
	}
	c.Private = NewPrivate(c)
	c.Public = NewPublic(c)
	c.Trade = NewTrade(c)
	return c
}

// Connect into the server
//
// https://www.okex.com/docs-v5/en/#websocket-api-connect
func (c *ClientWs) Connect(p bool) error {
	var connExists bool
	if p {
		c.PrivateConnMutex.RLock()
		connExists = c.PrivateConn != nil
		c.PrivateConnMutex.RUnlock()
	} else {
		c.PublicConnMutex.RLock()
		connExists = c.PublicConn != nil
		c.PublicConnMutex.RUnlock()
	}
	if connExists {
		return nil
	}
	err := c.dial(p)
	if err == nil {
		return nil
	}
	ticker := time.NewTicker(redialTick)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err = c.dial(p)
			if err == nil {
				return nil
			}
		case <-c.ctx.Done():
			return c.handleCancel("connect")
		}
	}
}

// Login
//
// https://www.okex.com/docs-v5/en/#websocket-api-login
func (c *ClientWs) Login() error {
	if c.Authorized {
		return nil
	}
	if c.AuthRequested != nil && time.Since(*c.AuthRequested).Seconds() < 30 {
		return nil
	}
	now := time.Now()
	c.AuthRequested = &now
	method := http.MethodGet
	path := "/users/self/verify"
	ts, sign := c.sign(method, path)
	args := []map[string]string{
		{
			"apiKey":     c.apiKey,
			"passphrase": c.passphrase,
			"timestamp":  ts,
			"sign":       sign,
		},
	}
	return c.Send(true, okex.LoginOperation, args)
}

// Subscribe
// Users can choose to subscribe to one or more channels, and the total length of multiple channels cannot exceed 4096 bytes.
//
// https://www.okex.com/docs-v5/en/#websocket-api-subscribe
func (c *ClientWs) Subscribe(p bool, ch []okex.ChannelName, args map[string]string) error {
	count := 1
	if len(ch) != 0 {
		count = len(ch)
	}
	tmpArgs := make([]map[string]string, count)
	tmpArgs[0] = args
	for i, name := range ch {
		tmpArgs[i] = map[string]string{}
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}
	return c.Send(p, okex.SubscribeOperation, tmpArgs)
}

// Unsubscribe into channel(s)
//
// https://www.okex.com/docs-v5/en/#websocket-api-unsubscribe
func (c *ClientWs) Unsubscribe(p bool, ch []okex.ChannelName, args map[string]string) error {
	tmpArgs := make([]map[string]string, len(ch))
	for i, name := range ch {
		tmpArgs[i] = make(map[string]string)
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}
	return c.Send(p, okex.UnsubscribeOperation, tmpArgs)
}

// Send message through either connections
func (c *ClientWs) Send(p bool, op okex.Operation, args []map[string]string, extras ...map[string]string) error {
	if op != okex.LoginOperation {
		err := c.Connect(p)
		if err == nil {
			if p {
				err = c.WaitForAuthorization()
				if err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}

	data := map[string]interface{}{
		"op":   op,
		"args": args,
	}
	for _, extra := range extras {
		for k, v := range extra {
			data[k] = v
		}
	}
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if p {
		c.PrivateSendChan <- j
	} else {
		c.PublicSendChan <- j
	}
	return nil
}

// SetChannels to receive certain events on separate channel
func (c *ClientWs) SetChannels(errCh chan *events.Error, subCh chan *events.Subscribe, unSub chan *events.Unsubscribe, lCh chan *events.Login, sCh chan *events.Success, wsErrCh chan error) {
	c.ErrChan = errCh
	c.SubscribeChan = subCh
	c.UnsubscribeCh = unSub
	c.LoginChan = lCh
	c.SuccessChan = sCh
	c.WSErrChan = wsErrCh
}

// WaitForAuthorization waits for the auth response and try to log in if it was needed
func (c *ClientWs) WaitForAuthorization() error {
	if c.Authorized {
		return nil
	}
	if err := c.Login(); err != nil {
		return err
	}
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()
	for range ticker.C {
		if c.Authorized {
			return nil
		}
	}
	return nil
}

func (c *ClientWs) Close(p bool) error {
	var conn *websocket.Conn
	if p {
		c.PrivateConnMutex.Lock()
		conn = c.PrivateConn
		defer c.PrivateConnMutex.Unlock()
	} else {
		c.PublicConnMutex.Lock()
		conn = c.PublicConn
		defer c.PublicConnMutex.Unlock()
	}
	if conn == nil {
		return nil
	}
	if p {
		c.PrivateConn = nil
	} else {
		c.PublicConn = nil
	}
	c.Authorized = false
	return conn.Close()
}

func (c *ClientWs) dial(p bool) error {
	conn, res, err := websocket.DefaultDialer.Dial(string(c.url[p]), nil)
	if err != nil {
		var statusCode int
		if res != nil {
			statusCode = res.StatusCode
		}
		return fmt.Errorf("error %d: %w", statusCode, err)
	}
	defer res.Body.Close()
	go func() {
		err := c.receiver(p)
		if err != nil {
			fmt.Printf("receiver error: %v\n", err)
			c.WSErrChan <- err
		}
	}()
	go func() {
		err := c.sender(p)
		if err != nil {
			fmt.Printf("sender error: %v\n", err)
			c.WSErrChan <- err
		}
	}()
	if p {
		c.PrivateConnMutex.Lock()
		c.PrivateConn = conn
		c.PrivateConnMutex.Unlock()
	} else {
		c.PublicConnMutex.Lock()
		c.PublicConn = conn
		c.PublicConnMutex.Unlock()
	}
	return nil
}

func (c *ClientWs) sendData(p bool, data []byte) error {
	var conn *websocket.Conn
	if p {
		c.PrivateConnMutex.RLock()
		conn = c.PrivateConn
		c.PrivateConnMutex.RUnlock()
	} else {
		c.PublicConnMutex.RLock()
		conn = c.PublicConn
		c.PublicConnMutex.RUnlock()
	}
	err := conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		return err
	}
	if p {
		c.PrivateWriteMutex.Lock()
	} else {
		c.PublicWriteMutex.Lock()
	}
	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		if p {
			c.PrivateWriteMutex.Unlock()
		} else {
			c.PublicWriteMutex.Unlock()
		}
		return err
	}
	if p {
		c.PrivateWriteMutex.Unlock()
	} else {
		c.PublicWriteMutex.Unlock()
	}
	now := time.Now()
	c.lastTransmit[p] = &now
	return nil
}

func (c *ClientWs) sender(p bool) error {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()
	for {
		select {
		case data := <-c.PrivateSendChan:
			c.sendData(true, data)
		case data := <-c.PublicSendChan:
			c.sendData(false, data)
		case <-ticker.C:
			var connExists bool
			if p {
				c.PrivateConnMutex.RLock()
				connExists = c.PrivateConn != nil
				c.PrivateConnMutex.RUnlock()
			} else {
				c.PublicConnMutex.RLock()
				connExists = c.PublicConn != nil
				c.PublicConnMutex.RUnlock()
			}

			if connExists && (c.lastTransmit[p] == nil || (c.lastTransmit[p] != nil && time.Since(*c.lastTransmit[p]) > PingPeriod)) {
				go func(p bool) {
					if p {
						c.PrivateSendChan <- []byte("ping")
					} else {
						c.PublicSendChan <- []byte("ping")
					}
				}(p)
			}
		case <-c.ctx.Done():
			return c.handleCancel("sender")
		}
	}
}

func (c *ClientWs) receiver(p bool) error {
	for {
		select {
		case <-c.ctx.Done():
			return c.handleCancel("receiver")
		default:
			var conn *websocket.Conn
			if p {
				c.PrivateConnMutex.RLock()
				conn = c.PrivateConn
			} else {
				c.PublicConnMutex.RLock()
				conn = c.PublicConn
			}
			err := conn.SetReadDeadline(time.Now().Add(pongWait))
			if err != nil {
				if p {
					c.PrivateConnMutex.RUnlock()
				} else {
					c.PublicConnMutex.RUnlock()
				}
				return err
			}
			mt, data, err := conn.ReadMessage()
			if err != nil {
				if p {
					c.PrivateConnMutex.RUnlock()
				} else {
					c.PublicConnMutex.RUnlock()
				}
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					return conn.Close()
				}
				return err
			}
			if p {
				c.PrivateConnMutex.RUnlock()
			} else {
				c.PublicConnMutex.RUnlock()
			}
			now := time.Now()
			c.lastTransmitMutex.Lock()
			c.lastTransmit[p] = &now
			c.lastTransmitMutex.Unlock()
			if mt == websocket.TextMessage && string(data) != "pong" {
				e := &events.Basic{}
				if err := json.Unmarshal(data, &e); err != nil {
					return err
				}
				go func() {
					c.process(data, e)
				}()
			}
		}
	}
}

func (c *ClientWs) sign(method, path string) (string, string) {
	t := time.Now().UTC().Unix()
	ts := fmt.Sprint(t)
	s := ts + method + path
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *ClientWs) handleCancel(msg string) error {
	go func() {
		c.DoneChan <- msg
	}()
	return fmt.Errorf("operation cancelled: %s", msg)
}

// TODO: break each case into a separate function
func (c *ClientWs) process(data []byte, e *events.Basic) bool {
	switch e.Event {
	case "error":
		e := events.Error{}
		_ = json.Unmarshal(data, &e)
		go func() {
			c.ErrChan <- &e
		}()
		return true
	case "subscribe":
		e := events.Subscribe{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.SubscribeChan != nil {
				c.SubscribeChan <- &e
			}
			c.StructuredEventChan <- e
		}()
		return true
	case "unsubscribe":
		e := events.Unsubscribe{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.UnsubscribeCh != nil {
				c.UnsubscribeCh <- &e
			}
			c.StructuredEventChan <- e
		}()
		return true
	case "login":
		if time.Since(*c.AuthRequested).Seconds() > 30 {
			c.AuthRequested = nil
			_ = c.Login()
			break
		}
		c.Authorized = true
		e := events.Login{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.LoginChan != nil {
				c.LoginChan <- &e
			}
			c.StructuredEventChan <- e
		}()
		return true
	}
	if c.Private.Process(data, e) {
		return true
	}
	if c.Public.Process(data, e) {
		return true
	}
	if e.ID != "" {
		if e.Code != 0 {
			ee := *e
			ee.Event = "error"
			return c.process(data, &ee)
		}
		e := events.Success{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.SuccessChan != nil {
				c.SuccessChan <- &e
			}
			c.StructuredEventChan <- e
		}()
		return true
	}
	go func() { c.RawEventChan <- e }()
	return false
}
