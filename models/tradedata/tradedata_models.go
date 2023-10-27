package tradedata

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/game-for-one/okex"
)

type (
	SupportCoin struct {
		Contract []string `json:"contract"`
		Option   []string `json:"option"`
		Spot     []string `json:"spot"`
	}
	TakerVolume struct {
		SellVol string
		BuyVol  string
		TS      okex.JSONTime
	}
	Ratio struct {
		Ratio string
		TS    okex.JSONTime
	}
	InterestAndVolumeRatio struct {
		Oi  string
		Vol string
		TS  okex.JSONTime
	}
	PutCallRatio struct {
		OiRatio  string
		VolRatio string
		TS       okex.JSONTime
	}
	InterestAndVolumeExpiry struct {
		CallOI  string
		PutOI   string
		CallVol string
		PutVol  string
		ExpTime okex.JSONTime
		TS      okex.JSONTime
	}
	InterestAndVolumeStrike struct {
		Strike  string
		CallOI  string
		PutOI   string
		CallVol string
		PutVol  string
		TS      okex.JSONTime
	}
	TakerFlow struct {
		CallBuyVol   string
		CallSellVol  string
		PutBuyVol    string
		PutSellVol   string
		CallBlockVol string
		PutBlockVol  string
		TS           okex.JSONTime
	}
)

func (c *TakerVolume) UnmarshalJSON(buf []byte) error {
	var (
		sellVol, buyVol, ts string
		err                 error
	)
	tmp := []interface{}{&ts, &sellVol, &buyVol}
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

	c.SellVol = sellVol
	c.BuyVol = buyVol

	return nil
}

func (c *Ratio) UnmarshalJSON(buf []byte) error {
	var (
		ratio, ts string
		err       error
	)
	tmp := []interface{}{&ts, &ratio}
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

	c.Ratio = ratio

	return nil
}

func (c *InterestAndVolumeRatio) UnmarshalJSON(buf []byte) error {
	var (
		oi, vol, ts string
		err         error
	)
	tmp := []interface{}{&ts, &oi, &vol}
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

	if oi != "" {
		c.Oi = oi
	}
	if vol != "" {
		c.Vol = vol
	}

	return nil
}

func (c *PutCallRatio) UnmarshalJSON(buf []byte) error {
	var (
		oi, vol, ts string
		err         error
	)
	tmp := []interface{}{&ts, &oi, &vol}
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

	if oi != "" {
		c.OiRatio = oi
	}
	if vol != "" {
		c.VolRatio = vol
	}

	return nil
}

func (c *InterestAndVolumeExpiry) UnmarshalJSON(buf []byte) error {
	var (
		callOI, putOI, callVol, putVol, expTime, ts string
		err                                         error
	)
	tmp := []interface{}{&ts, &expTime, &callOI, &putOI, &callVol, &putVol}
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

	exp, err := time.Parse("20060102", expTime)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.ExpTime) = exp

	if callOI != "" {
		c.CallOI = callOI
	}
	if putOI != "" {
		c.PutOI = putOI
	}
	if callVol != "" {
		c.CallVol = callVol
	}
	if putVol != "" {
		c.PutVol = putVol
	}

	return nil
}

func (c *InterestAndVolumeStrike) UnmarshalJSON(buf []byte) error {
	var (
		callOI, putOI, callVol, putVol, strike, ts string
		err                                        error
	)
	tmp := []interface{}{&ts, &strike, &callOI, &putOI, &callVol, &putVol}
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

	if callOI != "" {
		c.CallOI = callOI
	}
	if putOI != "" {
		c.PutOI = putOI
	}
	if callVol != "" {
		c.CallVol = callVol
	}
	if putVol != "" {
		c.PutVol = putVol
	}
	if strike != "" {
		c.Strike = strike
	}

	return nil
}

func (c *TakerFlow) UnmarshalJSON(buf []byte) error {
	var (
		ts, callBuyVol, callSellVol, putBuyVol, putSellVol, callBlockVol, putBlockVol string
		err                                                                           error
	)
	tmp := []interface{}{&ts, &callBuyVol, &callSellVol, &putBuyVol, &putSellVol, &callBlockVol, &putBlockVol}
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

	if callBuyVol != "" {
		c.CallBlockVol = callBuyVol
	}
	if callSellVol != "" {
		c.CallSellVol = callSellVol
	}
	if putBuyVol != "" {
		c.PutBuyVol = putBuyVol
	}
	if putSellVol != "" {
		c.PutSellVol = putSellVol
	}
	if callBlockVol != "" {
		c.CallBuyVol = callBlockVol
	}
	if putBlockVol != "" {
		c.PutBuyVol = putBlockVol
	}

	return nil
}
