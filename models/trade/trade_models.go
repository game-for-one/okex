package trade

import (
	"github.com/game-for-one/okex"
)

type (
	PlaceOrder struct {
		ClOrdID string          `json:"clOrdId"`
		Tag     string          `json:"tag"`
		SMsg    string          `json:"sMsg"`
		SCode   okex.JSONInt64  `json:"sCode"`
		OrdID   okex.JSONString `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string          `json:"ordId"`
		ClOrdID string          `json:"clOrdId"`
		SMsg    string          `json:"sMsg"`
		SCode   okex.JSONString `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string          `json:"ordId"`
		ClOrdID string          `json:"clOrdId"`
		ReqID   string          `json:"reqId"`
		SMsg    string          `json:"sMsg"`
		SCode   okex.JSONString `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID      string              `json:"instId"`
		Ccy         string              `json:"ccy"`
		OrdID       string              `json:"ordId"`
		ClOrdID     string              `json:"clOrdId"`
		TradeID     string              `json:"tradeId"`
		Tag         string              `json:"tag"`
		Category    string              `json:"category"`
		FeeCcy      string              `json:"feeCcy"`
		RebateCcy   string              `json:"rebateCcy"`
		Px          okex.JSONString     `json:"px"`
		Sz          okex.JSONString     `json:"sz"`
		Pnl         okex.JSONString     `json:"pnl"`
		AccFillSz   okex.JSONString     `json:"accFillSz"`
		FillPx      okex.JSONString     `json:"fillPx"`
		FillSz      okex.JSONString     `json:"fillSz"`
		FillTime    okex.JSONString     `json:"fillTime"`
		AvgPx       okex.JSONString     `json:"avgPx"`
		Lever       okex.JSONString     `json:"lever"`
		TpTriggerPx okex.JSONString     `json:"tpTriggerPx"`
		TpOrdPx     okex.JSONString     `json:"tpOrdPx"`
		SlTriggerPx okex.JSONString     `json:"slTriggerPx"`
		SlOrdPx     okex.JSONString     `json:"slOrdPx"`
		Fee         okex.JSONString     `json:"fee"`
		Rebate      okex.JSONString     `json:"rebate"`
		State       okex.OrderState     `json:"state"`
		TdMode      okex.TradeMode      `json:"tdMode"`
		PosSide     okex.PositionSide   `json:"posSide"`
		Side        okex.OrderSide      `json:"side"`
		OrdType     okex.OrderType      `json:"ordType"`
		InstType    okex.InstrumentType `json:"instType"`
		TgtCcy      okex.QuantityType   `json:"tgtCcy"`
		UTime       okex.JSONTime       `json:"uTime"`
		CTime       okex.JSONTime       `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string              `json:"instId"`
		OrdID    string              `json:"ordId"`
		TradeID  string              `json:"tradeId"`
		ClOrdID  string              `json:"clOrdId"`
		BillID   string              `json:"billId"`
		Tag      okex.JSONString     `json:"tag"`
		FillPx   okex.JSONString     `json:"fillPx"`
		FillSz   okex.JSONString     `json:"fillSz"`
		FeeCcy   okex.JSONString     `json:"feeCcy"`
		Fee      okex.JSONString     `json:"fee"`
		InstType okex.InstrumentType `json:"instType"`
		Side     okex.OrderSide      `json:"side"`
		PosSide  okex.PositionSide   `json:"posSide"`
		ExecType okex.OrderFlowType  `json:"execType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PlaceAlgoOrder struct {
		AlgoID string         `json:"algoId"`
		SMsg   string         `json:"sMsg"`
		SCode  okex.JSONInt64 `json:"sCode"`
	}
	CancelAlgoOrder struct {
		AlgoID string         `json:"algoId"`
		SMsg   string         `json:"sMsg"`
		SCode  okex.JSONInt64 `json:"sCode"`
	}
	AlgoOrder struct {
		InstID       string              `json:"instId"`
		Ccy          string              `json:"ccy"`
		OrdID        string              `json:"ordId"`
		AlgoID       string              `json:"algoId"`
		ClOrdID      string              `json:"clOrdId"`
		TradeID      string              `json:"tradeId"`
		Tag          string              `json:"tag"`
		Category     string              `json:"category"`
		FeeCcy       string              `json:"feeCcy"`
		RebateCcy    string              `json:"rebateCcy"`
		TimeInterval string              `json:"timeInterval"`
		Px           okex.JSONString     `json:"px"`
		PxVar        okex.JSONString     `json:"pxVar"`
		PxSpread     okex.JSONString     `json:"pxSpread"`
		PxLimit      okex.JSONString     `json:"pxLimit"`
		Sz           okex.JSONString     `json:"sz"`
		SzLimit      okex.JSONString     `json:"szLimit"`
		ActualSz     okex.JSONString     `json:"actualSz"`
		ActualPx     okex.JSONString     `json:"actualPx"`
		Pnl          okex.JSONString     `json:"pnl"`
		AccFillSz    okex.JSONString     `json:"accFillSz"`
		FillPx       okex.JSONString     `json:"fillPx"`
		FillSz       okex.JSONString     `json:"fillSz"`
		FillTime     okex.JSONString     `json:"fillTime"`
		AvgPx        okex.JSONString     `json:"avgPx"`
		Lever        okex.JSONString     `json:"lever"`
		TpTriggerPx  okex.JSONString     `json:"tpTriggerPx"`
		TpOrdPx      okex.JSONString     `json:"tpOrdPx"`
		SlTriggerPx  okex.JSONString     `json:"slTriggerPx"`
		SlOrdPx      okex.JSONString     `json:"slOrdPx"`
		OrdPx        okex.JSONString     `json:"ordPx"`
		Fee          okex.JSONString     `json:"fee"`
		Rebate       okex.JSONString     `json:"rebate"`
		State        okex.OrderState     `json:"state"`
		TdMode       okex.TradeMode      `json:"tdMode"`
		ActualSide   okex.PositionSide   `json:"actualSide"`
		PosSide      okex.PositionSide   `json:"posSide"`
		Side         okex.OrderSide      `json:"side"`
		OrdType      okex.AlgoOrderType  `json:"ordType"`
		InstType     okex.InstrumentType `json:"instType"`
		TgtCcy       okex.QuantityType   `json:"tgtCcy"`
		CTime        okex.JSONTime       `json:"cTime"`
		TriggerTime  okex.JSONTime       `json:"triggerTime"`
	}
)
