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
		OrdID   okex.JSONAmount `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string          `json:"ordId"`
		ClOrdID string          `json:"clOrdId"`
		SMsg    string          `json:"sMsg"`
		SCode   okex.JSONAmount `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string          `json:"ordId"`
		ClOrdID string          `json:"clOrdId"`
		ReqID   string          `json:"reqId"`
		SMsg    string          `json:"sMsg"`
		SCode   okex.JSONAmount `json:"sCode"`
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
		Px          okex.JSONAmount     `json:"px"`
		Sz          okex.JSONAmount     `json:"sz"`
		Pnl         okex.JSONAmount     `json:"pnl"`
		AccFillSz   okex.JSONAmount     `json:"accFillSz"`
		FillPx      okex.JSONAmount     `json:"fillPx"`
		FillSz      okex.JSONAmount     `json:"fillSz"`
		FillTime    okex.JSONAmount     `json:"fillTime"`
		AvgPx       okex.JSONAmount     `json:"avgPx"`
		Lever       okex.JSONAmount     `json:"lever"`
		TpTriggerPx okex.JSONAmount     `json:"tpTriggerPx"`
		TpOrdPx     okex.JSONAmount     `json:"tpOrdPx"`
		SlTriggerPx okex.JSONAmount     `json:"slTriggerPx"`
		SlOrdPx     okex.JSONAmount     `json:"slOrdPx"`
		Fee         okex.JSONAmount     `json:"fee"`
		Rebate      okex.JSONAmount     `json:"rebate"`
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
		Tag      okex.JSONAmount     `json:"tag"`
		FillPx   okex.JSONAmount     `json:"fillPx"`
		FillSz   okex.JSONAmount     `json:"fillSz"`
		FeeCcy   okex.JSONAmount     `json:"feeCcy"`
		Fee      okex.JSONAmount     `json:"fee"`
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
		Px           okex.JSONAmount     `json:"px"`
		PxVar        okex.JSONAmount     `json:"pxVar"`
		PxSpread     okex.JSONAmount     `json:"pxSpread"`
		PxLimit      okex.JSONAmount     `json:"pxLimit"`
		Sz           okex.JSONAmount     `json:"sz"`
		SzLimit      okex.JSONAmount     `json:"szLimit"`
		ActualSz     okex.JSONAmount     `json:"actualSz"`
		ActualPx     okex.JSONAmount     `json:"actualPx"`
		Pnl          okex.JSONAmount     `json:"pnl"`
		AccFillSz    okex.JSONAmount     `json:"accFillSz"`
		FillPx       okex.JSONAmount     `json:"fillPx"`
		FillSz       okex.JSONAmount     `json:"fillSz"`
		FillTime     okex.JSONAmount     `json:"fillTime"`
		AvgPx        okex.JSONAmount     `json:"avgPx"`
		Lever        okex.JSONAmount     `json:"lever"`
		TpTriggerPx  okex.JSONAmount     `json:"tpTriggerPx"`
		TpOrdPx      okex.JSONAmount     `json:"tpOrdPx"`
		SlTriggerPx  okex.JSONAmount     `json:"slTriggerPx"`
		SlOrdPx      okex.JSONAmount     `json:"slOrdPx"`
		OrdPx        okex.JSONAmount     `json:"ordPx"`
		Fee          okex.JSONAmount     `json:"fee"`
		Rebate       okex.JSONAmount     `json:"rebate"`
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
