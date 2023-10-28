package market

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/game-for-one/okex"
)

type (
	Ticker struct {
		InstID    string              `json:"instId"`
		Last      string              `json:"last"`
		LastSz    string              `json:"lastSz"`
		AskPx     string              `json:"askPx"`
		AskSz     string              `json:"askSz"`
		BidPx     string              `json:"bidPx"`
		BidSz     string              `json:"bidSz"`
		Open24h   string              `json:"open24h"`
		High24h   string              `json:"high24h"`
		Low24h    string              `json:"low24h"`
		VolCcy24h string              `json:"volCcy24h"`
		Vol24h    string              `json:"vol24h"`
		SodUtc0   string              `json:"sodUtc0"`
		SodUtc8   string              `json:"sodUtc8"`
		InstType  okex.InstrumentType `json:"instType"`
		TS        okex.JSONTime       `json:"ts"`
	}
	IndexTicker struct {
		InstID  string        `json:"instId"`
		IdxPx   string        `json:"idxPx"`
		High24h string        `json:"high24h"`
		Low24h  string        `json:"low24h"`
		Open24h string        `json:"open24h"`
		SodUtc0 string        `json:"sodUtc0"`
		SodUtc8 string        `json:"sodUtc8"`
		TS      okex.JSONTime `json:"ts"`
	}
	OrderBook struct {
		Asks []*OrderBookEntity `json:"asks"`
		Bids []*OrderBookEntity `json:"bids"`
		TS   okex.JSONTime      `json:"ts"`
	}
	OrderBookWs struct {
		Asks     []*OrderBookEntity `json:"asks"`
		Bids     []*OrderBookEntity `json:"bids"`
		Checksum int                `json:"checksum"`
		TS       okex.JSONTime      `json:"ts"`
	}
	OrderBookEntity struct {
		DepthPrice      string
		Size            string
		LiquidatedOrder int
		OrderNumbers    int
	}
	Candle struct {
		O      string
		H      string
		L      string
		C      string
		Vol    string
		VolCcy string
		TS     okex.JSONTime
	}
	IndexCandle struct {
		O  string
		H  string
		L  string
		C  string
		TS okex.JSONTime
	}
	Trade struct {
		InstID  string         `json:"instId"`
		TradeID string         `json:"tradeId"`
		Px      string         `json:"px"`
		Sz      string         `json:"sz"`
		Side    okex.TradeSide `json:"side"`
		TS      okex.JSONTime  `json:"ts"`
	}
	TotalVolume24H struct {
		VolUsd string        `json:"volUsd"`
		VolCny string        `json:"volCny"`
		TS     okex.JSONTime `json:"ts"`
	}
	IndexComponent struct {
		Index      string        `json:"index"`
		Last       string        `json:"last"`
		Components []*Component  `json:"components"`
		TS         okex.JSONTime `json:"ts"`
	}
	Component struct {
		Exch   string `json:"exch"`
		Symbol string `json:"symbol"`
		SymPx  string `json:"symPx"`
		Wgt    string `json:"wgt"`
		CnvPx  string `json:"cnvPx"`
	}
)

func (o *OrderBookEntity) UnmarshalJSON(buf []byte) error {
	var (
		dp, s, lo, on string
		err           error
	)
	tmp := []interface{}{&dp, &s, &lo, &on}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in OrderBookEntity: %d != %d", g, e)
	}
	o.DepthPrice = dp
	o.Size = s
	o.LiquidatedOrder, err = strconv.Atoi(lo)
	if err != nil {
		return err
	}
	o.OrderNumbers, err = strconv.Atoi(on)
	if err != nil {
		return err
	}

	return nil
}

func (c *Candle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, vol, volCcy, ts string
		err                          error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl, &vol, &volCcy}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O = o
	c.H = h
	c.L = l
	c.C = cl
	c.Vol = vol
	c.VolCcy = volCcy

	return nil
}

func (c *IndexCandle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, ts string
		err             error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O = o
	c.H = h
	c.L = l
	c.C = cl

	return nil
}
