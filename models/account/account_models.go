package account

import (
	"github.com/game-for-one/okex"
)

type (
	Balance struct {
		TotalEq     okex.JSONString   `json:"totalEq"`
		IsoEq       okex.JSONString   `json:"isoEq"`
		AdjEq       okex.JSONString   `json:"adjEq,omitempty"`
		OrdFroz     okex.JSONString   `json:"ordFroz,omitempty"`
		Imr         okex.JSONString   `json:"imr,omitempty"`
		Mmr         okex.JSONString   `json:"mmr,omitempty"`
		MgnRatio    okex.JSONString   `json:"mgnRatio,omitempty"`
		NotionalUsd okex.JSONString   `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       okex.JSONTime     `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string          `json:"ccy"`
		Eq            okex.JSONString `json:"eq"`
		CashBal       okex.JSONString `json:"cashBal"`
		IsoEq         okex.JSONString `json:"isoEq,omitempty"`
		AvailEq       okex.JSONString `json:"availEq,omitempty"`
		DisEq         okex.JSONString `json:"disEq"`
		AvailBal      okex.JSONString `json:"availBal"`
		FrozenBal     okex.JSONString `json:"frozenBal"`
		OrdFrozen     okex.JSONString `json:"ordFrozen"`
		Liab          okex.JSONString `json:"liab,omitempty"`
		Upl           okex.JSONString `json:"upl,omitempty"`
		UplLib        okex.JSONString `json:"uplLib,omitempty"`
		CrossLiab     okex.JSONString `json:"crossLiab,omitempty"`
		IsoLiab       okex.JSONString `json:"isoLiab,omitempty"`
		MgnRatio      okex.JSONString `json:"mgnRatio,omitempty"`
		Interest      okex.JSONString `json:"interest,omitempty"`
		Twap          okex.JSONString `json:"twap,omitempty"`
		MaxLoan       okex.JSONString `json:"maxLoan,omitempty"`
		EqUsd         okex.JSONString `json:"eqUsd"`
		NotionalLever okex.JSONString `json:"notionalLever,omitempty"`
		StgyEq        okex.JSONString `json:"stgyEq"`
		IsoUpl        okex.JSONString `json:"isoUpl,omitempty"`
		UTime         okex.JSONTime   `json:"uTime"`
	}
	Position struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		LiabCcy     string              `json:"liabCcy,omitempty"`
		OptVal      string              `json:"optVal,omitempty"`
		Ccy         string              `json:"ccy"`
		PosID       string              `json:"posId"`
		TradeID     string              `json:"tradeId"`
		Pos         okex.JSONString     `json:"pos"`
		AvailPos    okex.JSONString     `json:"availPos,omitempty"`
		AvgPx       okex.JSONString     `json:"avgPx"`
		Upl         okex.JSONString     `json:"upl"`
		UplRatio    okex.JSONString     `json:"uplRatio"`
		Lever       okex.JSONString     `json:"lever"`
		LiqPx       okex.JSONString     `json:"liqPx,omitempty"`
		Imr         okex.JSONString     `json:"imr,omitempty"`
		Margin      okex.JSONString     `json:"margin,omitempty"`
		MgnRatio    okex.JSONString     `json:"mgnRatio"`
		Mmr         okex.JSONString     `json:"mmr"`
		Liab        okex.JSONString     `json:"liab,omitempty"`
		Interest    okex.JSONString     `json:"interest"`
		NotionalUsd okex.JSONString     `json:"notionalUsd"`
		ADL         okex.JSONString     `json:"adl"`
		Last        okex.JSONString     `json:"last"`
		DeltaBS     okex.JSONString     `json:"deltaBS"`
		DeltaPA     okex.JSONString     `json:"deltaPA"`
		GammaBS     okex.JSONString     `json:"gammaBS"`
		GammaPA     okex.JSONString     `json:"gammaPA"`
		ThetaBS     okex.JSONString     `json:"thetaBS"`
		ThetaPA     okex.JSONString     `json:"thetaPA"`
		VegaBS      okex.JSONString     `json:"vegaBS"`
		VegaPA      okex.JSONString     `json:"vegaPA"`
		PosSide     okex.PositionSide   `json:"posSide"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
		InstType    okex.InstrumentType `json:"instType"`
		CTime       okex.JSONTime       `json:"cTime"`
		UTime       okex.JSONTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okex.EventType    `json:"eventType"`
		PTime     okex.JSONTime     `json:"pTime"`
		UTime     okex.JSONTime     `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   okex.JSONString                      `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okex.JSONTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string          `json:"ccy"`
		Eq    okex.JSONString `json:"eq"`
		DisEq okex.JSONString `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		Ccy         string              `json:"ccy"`
		NotionalCcy okex.JSONString     `json:"notionalCcy"`
		Pos         okex.JSONString     `json:"pos"`
		NotionalUsd okex.JSONString     `json:"notionalUsd"`
		PosSide     okex.PositionSide   `json:"posSide"`
		InstType    okex.InstrumentType `json:"instType"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string              `json:"ccy"`
		InstID    string              `json:"instId"`
		Notes     string              `json:"notes"`
		BillID    string              `json:"billId"`
		OrdID     string              `json:"ordId"`
		BalChg    okex.JSONString     `json:"balChg"`
		PosBalChg okex.JSONString     `json:"posBalChg"`
		Bal       okex.JSONString     `json:"bal"`
		PosBal    okex.JSONString     `json:"posBal"`
		Sz        okex.JSONString     `json:"sz"`
		Pnl       okex.JSONString     `json:"pnl"`
		Fee       okex.JSONString     `json:"fee"`
		From      okex.AccountType    `json:"from,string"`
		To        okex.AccountType    `json:"to,string"`
		InstType  okex.InstrumentType `json:"instType"`
		MgnMode   okex.MarginMode     `json:"MgnMode"`
		Type      okex.BillType       `json:"type,string"`
		SubType   okex.BillSubType    `json:"subType,string"`
		TS        okex.JSONTime       `json:"ts"`
	}
	Config struct {
		Level      string            `json:"level"`
		LevelTmp   string            `json:"levelTmp"`
		AcctLv     string            `json:"acctLv"`
		AutoLoan   bool              `json:"autoLoan"`
		UID        string            `json:"uid"`
		GreeksType okex.GreekType    `json:"greeksType"`
		PosMode    okex.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode okex.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string            `json:"instId"`
		Lever   okex.JSONString   `json:"lever"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string          `json:"instId"`
		Ccy     string          `json:"ccy"`
		MaxBuy  okex.JSONString `json:"maxBuy"`
		MaxSell okex.JSONString `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string          `json:"instId"`
		AvailBuy  okex.JSONString `json:"availBuy"`
		AvailSell okex.JSONString `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string            `json:"instId"`
		Amt     okex.JSONString   `json:"amt"`
		PosSide okex.PositionSide `json:"posSide,string"`
		Type    okex.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string          `json:"instId"`
		MgnCcy  string          `json:"mgnCcy"`
		Ccy     string          `json:"ccy"`
		MaxLoan okex.JSONString `json:"maxLoan"`
		MgnMode okex.MarginMode `json:"mgnMode"`
		Side    okex.OrderSide  `json:"side,string"`
	}
	Fee struct {
		Level    string              `json:"level"`
		Taker    okex.JSONString     `json:"taker"`
		Maker    okex.JSONString     `json:"maker"`
		Delivery okex.JSONString     `json:"delivery,omitempty"`
		Exercise okex.JSONString     `json:"exercise,omitempty"`
		Category okex.FeeCategory    `json:"category,string"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string          `json:"instId"`
		Ccy          string          `json:"ccy"`
		Interest     okex.JSONString `json:"interest"`
		InterestRate okex.JSONString `json:"interestRate"`
		Liab         okex.JSONString `json:"liab"`
		MgnMode      okex.MarginMode `json:"mgnMode"`
		TS           okex.JSONTime   `json:"ts"`
	}
	InterestRate struct {
		Ccy          string          `json:"ccy"`
		InterestRate okex.JSONString `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string          `json:"ccy"`
		MaxWd okex.JSONString `json:"maxWd"`
	}
)
