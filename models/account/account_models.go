package account

import (
	"github.com/game-for-one/okex"
)

type (
	Balance struct {
		TotalEq     okex.JSONAmount   `json:"totalEq"`
		IsoEq       okex.JSONAmount   `json:"isoEq"`
		AdjEq       okex.JSONAmount   `json:"adjEq,omitempty"`
		OrdFroz     okex.JSONAmount   `json:"ordFroz,omitempty"`
		Imr         okex.JSONAmount   `json:"imr,omitempty"`
		Mmr         okex.JSONAmount   `json:"mmr,omitempty"`
		MgnRatio    okex.JSONAmount   `json:"mgnRatio,omitempty"`
		NotionalUsd okex.JSONAmount   `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       okex.JSONTime     `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string          `json:"ccy"`
		Eq            okex.JSONAmount `json:"eq"`
		CashBal       okex.JSONAmount `json:"cashBal"`
		IsoEq         okex.JSONAmount `json:"isoEq,omitempty"`
		AvailEq       okex.JSONAmount `json:"availEq,omitempty"`
		DisEq         okex.JSONAmount `json:"disEq"`
		AvailBal      okex.JSONAmount `json:"availBal"`
		FrozenBal     okex.JSONAmount `json:"frozenBal"`
		OrdFrozen     okex.JSONAmount `json:"ordFrozen"`
		Liab          okex.JSONAmount `json:"liab,omitempty"`
		Upl           okex.JSONAmount `json:"upl,omitempty"`
		UplLib        okex.JSONAmount `json:"uplLib,omitempty"`
		CrossLiab     okex.JSONAmount `json:"crossLiab,omitempty"`
		IsoLiab       okex.JSONAmount `json:"isoLiab,omitempty"`
		MgnRatio      okex.JSONAmount `json:"mgnRatio,omitempty"`
		Interest      okex.JSONAmount `json:"interest,omitempty"`
		Twap          okex.JSONAmount `json:"twap,omitempty"`
		MaxLoan       okex.JSONAmount `json:"maxLoan,omitempty"`
		EqUsd         okex.JSONAmount `json:"eqUsd"`
		NotionalLever okex.JSONAmount `json:"notionalLever,omitempty"`
		StgyEq        okex.JSONAmount `json:"stgyEq"`
		IsoUpl        okex.JSONAmount `json:"isoUpl,omitempty"`
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
		Pos         okex.JSONAmount     `json:"pos"`
		AvailPos    okex.JSONAmount     `json:"availPos,omitempty"`
		AvgPx       okex.JSONAmount     `json:"avgPx"`
		Upl         okex.JSONAmount     `json:"upl"`
		UplRatio    okex.JSONAmount     `json:"uplRatio"`
		Lever       okex.JSONAmount     `json:"lever"`
		LiqPx       okex.JSONAmount     `json:"liqPx,omitempty"`
		Imr         okex.JSONAmount     `json:"imr,omitempty"`
		Margin      okex.JSONAmount     `json:"margin,omitempty"`
		MgnRatio    okex.JSONAmount     `json:"mgnRatio"`
		Mmr         okex.JSONAmount     `json:"mmr"`
		Liab        okex.JSONAmount     `json:"liab,omitempty"`
		Interest    okex.JSONAmount     `json:"interest"`
		NotionalUsd okex.JSONAmount     `json:"notionalUsd"`
		ADL         okex.JSONAmount     `json:"adl"`
		Last        okex.JSONAmount     `json:"last"`
		DeltaBS     okex.JSONAmount     `json:"deltaBS"`
		DeltaPA     okex.JSONAmount     `json:"deltaPA"`
		GammaBS     okex.JSONAmount     `json:"gammaBS"`
		GammaPA     okex.JSONAmount     `json:"gammaPA"`
		ThetaBS     okex.JSONAmount     `json:"thetaBS"`
		ThetaPA     okex.JSONAmount     `json:"thetaPA"`
		VegaBS      okex.JSONAmount     `json:"vegaBS"`
		VegaPA      okex.JSONAmount     `json:"vegaPA"`
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
		AdjEq   okex.JSONAmount                      `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okex.JSONTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string          `json:"ccy"`
		Eq    okex.JSONAmount `json:"eq"`
		DisEq okex.JSONAmount `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		Ccy         string              `json:"ccy"`
		NotionalCcy okex.JSONAmount     `json:"notionalCcy"`
		Pos         okex.JSONAmount     `json:"pos"`
		NotionalUsd okex.JSONAmount     `json:"notionalUsd"`
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
		BalChg    okex.JSONAmount     `json:"balChg"`
		PosBalChg okex.JSONAmount     `json:"posBalChg"`
		Bal       okex.JSONAmount     `json:"bal"`
		PosBal    okex.JSONAmount     `json:"posBal"`
		Sz        okex.JSONAmount     `json:"sz"`
		Pnl       okex.JSONAmount     `json:"pnl"`
		Fee       okex.JSONAmount     `json:"fee"`
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
		Lever   okex.JSONAmount   `json:"lever"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string          `json:"instId"`
		Ccy     string          `json:"ccy"`
		MaxBuy  okex.JSONAmount `json:"maxBuy"`
		MaxSell okex.JSONAmount `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string          `json:"instId"`
		AvailBuy  okex.JSONAmount `json:"availBuy"`
		AvailSell okex.JSONAmount `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string            `json:"instId"`
		Amt     okex.JSONAmount   `json:"amt"`
		PosSide okex.PositionSide `json:"posSide,string"`
		Type    okex.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string          `json:"instId"`
		MgnCcy  string          `json:"mgnCcy"`
		Ccy     string          `json:"ccy"`
		MaxLoan okex.JSONAmount `json:"maxLoan"`
		MgnMode okex.MarginMode `json:"mgnMode"`
		Side    okex.OrderSide  `json:"side,string"`
	}
	Fee struct {
		Level    string              `json:"level"`
		Taker    okex.JSONAmount     `json:"taker"`
		Maker    okex.JSONAmount     `json:"maker"`
		Delivery okex.JSONAmount     `json:"delivery,omitempty"`
		Exercise okex.JSONAmount     `json:"exercise,omitempty"`
		Category okex.FeeCategory    `json:"category,string"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string          `json:"instId"`
		Ccy          string          `json:"ccy"`
		Interest     okex.JSONAmount `json:"interest"`
		InterestRate okex.JSONAmount `json:"interestRate"`
		Liab         okex.JSONAmount `json:"liab"`
		MgnMode      okex.MarginMode `json:"mgnMode"`
		TS           okex.JSONTime   `json:"ts"`
	}
	InterestRate struct {
		Ccy          string          `json:"ccy"`
		InterestRate okex.JSONAmount `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string          `json:"ccy"`
		MaxWd okex.JSONAmount `json:"maxWd"`
	}
)
