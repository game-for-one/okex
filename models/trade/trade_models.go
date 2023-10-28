package trade

import (
	"github.com/game-for-one/okex"
)

type (
	PlaceOrder struct {
		ClOrdID string         `json:"clOrdId"`
		Tag     string         `json:"tag"`
		SMsg    string         `json:"sMsg"`
		SCode   okex.JSONInt64 `json:"sCode"`
		OrdID   string         `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string `json:"ordId"`
		ClOrdID string `json:"clOrdId"`
		SMsg    string `json:"sMsg"`
		SCode   string `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string `json:"ordId"`
		ClOrdID string `json:"clOrdId"`
		ReqID   string `json:"reqId"`
		SMsg    string `json:"sMsg"`
		SCode   string `json:"sCode"`
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
		Px          string              `json:"px"`
		Sz          string              `json:"sz"`
		Pnl         string              `json:"pnl"`
		AccFillSz   string              `json:"accFillSz"`
		FillPx      string              `json:"fillPx"`
		FillSz      string              `json:"fillSz"`
		FillTime    string              `json:"fillTime"`
		AvgPx       string              `json:"avgPx"`
		Lever       string              `json:"lever"`
		TpTriggerPx string              `json:"tpTriggerPx"`
		TpOrdPx     string              `json:"tpOrdPx"`
		SlTriggerPx string              `json:"slTriggerPx"`
		SlOrdPx     string              `json:"slOrdPx"`
		Fee         string              `json:"fee"`
		Rebate      string              `json:"rebate"`
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
		Tag      string              `json:"tag"`
		FillPx   string              `json:"fillPx"`
		FillSz   string              `json:"fillSz"`
		FeeCcy   string              `json:"feeCcy"`
		Fee      string              `json:"fee"`
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
		Px           string              `json:"px"`
		PxVar        string              `json:"pxVar"`
		PxSpread     string              `json:"pxSpread"`
		PxLimit      string              `json:"pxLimit"`
		Sz           string              `json:"sz"`
		SzLimit      string              `json:"szLimit"`
		ActualSz     string              `json:"actualSz"`
		ActualPx     string              `json:"actualPx"`
		Pnl          string              `json:"pnl"`
		AccFillSz    string              `json:"accFillSz"`
		FillPx       string              `json:"fillPx"`
		FillSz       string              `json:"fillSz"`
		FillTime     string              `json:"fillTime"`
		AvgPx        string              `json:"avgPx"`
		Lever        string              `json:"lever"`
		TpTriggerPx  string              `json:"tpTriggerPx"`
		TpOrdPx      string              `json:"tpOrdPx"`
		SlTriggerPx  string              `json:"slTriggerPx"`
		SlOrdPx      string              `json:"slOrdPx"`
		OrdPx        string              `json:"ordPx"`
		Fee          string              `json:"fee"`
		Rebate       string              `json:"rebate"`
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
