package trade

import (
	"github.com/game-for-one/okex"
)

type (
	PlaceOrder struct {
		ID         string            `json:"-"`
		InstID     string            `json:"instId"`
		Ccy        string            `json:"ccy,omitempty"`
		ClOrdID    string            `json:"clOrdId,omitempty"`
		Tag        string            `json:"tag,omitempty"`
		ReduceOnly bool              `json:"reduceOnly,omitempty"`
		Sz         string            `json:"sz,string"`
		Px         string            `json:"px,omitempty,string"`
		TdMode     okex.TradeMode    `json:"tdMode"`
		Side       okex.OrderSide    `json:"side"`
		PosSide    okex.PositionSide `json:"posSide,omitempty"`
		OrdType    okex.OrderType    `json:"ordType"`
		TgtCcy     okex.QuantityType `json:"tgtCcy,omitempty"`
	}
	CancelOrder struct {
		ID      string `json:"-"`
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	AmendOrder struct {
		ID        string `json:"-"`
		InstID    string `json:"instId"`
		OrdID     string `json:"ordId,omitempty"`
		ClOrdID   string `json:"clOrdId,omitempty"`
		ReqID     string `json:"reqId,omitempty"`
		NewSz     int64  `json:"newSz,omitempty,string"`
		NewPx     string `json:"newPx,omitempty,string"`
		CxlOnFail bool   `json:"cxlOnFail,omitempty"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId"`
		Ccy     string            `json:"ccy,omitempty"`
		PosSide okex.PositionSide `json:"posSide,omitempty"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
	}
	OrderDetails struct {
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	OrderList struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		After    string              `json:"after,omitempty,string"`
		Before   string              `json:"before,omitempty,string"`
		Limit    string              `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
		OrdType  okex.OrderType      `json:"ordType,omitempty"`
		State    okex.OrderState     `json:"state,omitempty"`
	}
	TransactionDetails struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		OrdID    string              `json:"ordId,omitempty"`
		After    string              `json:"after,omitempty,string"`
		Before   string              `json:"before,omitempty,string"`
		Limit    string              `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
	}
	PlaceAlgoOrder struct {
		InstID        string             `json:"instId"`
		TdMode        okex.TradeMode     `json:"tdMode"`
		Ccy           string             `json:"ccy,omitempty"`
		Side          okex.OrderSide     `json:"side"`
		PosSide       okex.PositionSide  `json:"posSide,omitempty"`
		OrdType       okex.AlgoOrderType `json:"ordType"`
		Sz            string             `json:"sz,omitempty"` // either this or closeFraction is required
		ReduceOnly    bool               `json:"reduceOnly,omitempty"`
		TgtCcy        okex.QuantityType  `json:"tgtCcy,omitempty"`        // Only applicable to SPOT traded with Market buy conditional order. Default is quote_ccy for buy, base_ccy for sell
		CloseFraction string             `json:"closeFraction,omitempty"` // either this or sz is required
		StopOrder
		TriggerOrder
		IcebergOrder
		TWAPOrder
		TrailingStopOrder
	}
	StopOrder struct {
		TpTriggerPx string `json:"tpTriggerPx,string,omitempty"`
		TpOrdPx     string `json:"tpOrdPx,string,omitempty"`
		SlTriggerPx string `json:"slTriggerPx,string,omitempty"`
		SlOrdPx     string `json:"slOrdPx,string,omitempty"`
	}
	TriggerOrder struct {
		TriggerPx string `json:"triggerPx,string,omitempty"`
		OrdPx     string `json:"ordPx,string,omitempty"`
	}
	IcebergOrder struct {
		PxVar    string `json:"pxVar,string,omitempty"`
		PxSpread string `json:"pxSpread,string,omitempty"`
		SzLimit  int64  `json:"szLimit,string"`
		PxLimit  string `json:"pxLimit"`
	}
	TWAPOrder struct {
		IcebergOrder
		TimeInterval string `json:"timeInterval"`
	}
	TrailingStopOrder struct {
		PercentRetracement string `json:"callbackRatio"` // 0.01 represents 1%
	}
	CancelAlgoOrder struct {
		InstID string `json:"instId"`
		AlgoID string `json:"AlgoId"`
	}
	AlgoOrderList struct {
		InstType okex.InstrumentType `json:"instType,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		After    string              `json:"after,omitempty,string"`
		Before   string              `json:"before,omitempty,string"`
		Limit    string              `json:"limit,omitempty,string"`
		OrdType  okex.AlgoOrderType  `json:"ordType,omitempty"`
		State    okex.OrderState     `json:"state,omitempty"`
	}
)
