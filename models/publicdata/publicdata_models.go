package publicdata

import (
	"github.com/game-for-one/okex"
)

type (
	Instrument struct {
		InstType   okex.InstrumentType `json:"instType"`
		InstID     string              `json:"instId"`
		InstFamily string              `json:"instFamily,omitempty"`

		Uly       string               `json:"uly,omitempty"`
		BaseCcy   string               `json:"baseCcy,omitempty"`
		QuoteCcy  string               `json:"quoteCcy,omitempty"`
		SettleCcy string               `json:"settleCcy,omitempty"`
		CtValCcy  string               `json:"ctValCcy,omitempty"`
		CtVal     okex.JSONString      `json:"ctVal,omitempty"`
		CtMult    okex.JSONString      `json:"ctMult,omitempty"`
		Stk       okex.JSONString      `json:"stk,omitempty"`
		TickSz    okex.JSONString      `json:"tickSz,omitempty"`
		LotSz     okex.JSONString      `json:"lotSz,omitempty"`
		MinSz     okex.JSONString      `json:"minSz,omitempty"`
		Lever     okex.JSONString      `json:"lever"`
		OptType   okex.OptionType      `json:"optType,omitempty"`
		ListTime  okex.JSONTime        `json:"listTime"`
		ExpTime   okex.JSONTime        `json:"expTime,omitempty"`
		CtType    okex.ContractType    `json:"ctType,omitempty"`
		Alias     okex.AliasType       `json:"alias,omitempty"`
		State     okex.InstrumentState `json:"state"`

		MaxLmSz      okex.JSONString `json:"maxLmtSz,omitempty"`
		MaxMkSz      okex.JSONString `json:"maxMktSz,omitempty"`
		MaxTwapSz    okex.JSONString `json:"maxTwapSz,omitempty"`
		MaxIcebergSz okex.JSONString `json:"maxIcebergSz,omitempty"`
		MaxTriggerSz okex.JSONString `json:"maxTriggerSz,omitempty"`
		MaxStopSz    okex.JSONString `json:"maxStopSz,omitempty"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okex.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId"`
		Px     okex.JSONString           `json:"px"`
		Type   okex.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string              `json:"instId"`
		Oi       okex.JSONString     `json:"oi"`
		OiCcy    okex.JSONString     `json:"oiCcy"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string              `json:"instId"`
		InstType        okex.InstrumentType `json:"instType"`
		FundingRate     okex.JSONString     `json:"fundingRate"`
		NextFundingRate okex.JSONString     `json:"NextFundingRate"`
		FundingTime     okex.JSONTime       `json:"fundingTime"`
		NextFundingTime okex.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		BuyLmt   okex.JSONString     `json:"buyLmt"`
		SellLmt  okex.JSONString     `json:"sellLmt"`
		TS       okex.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		SettlePx okex.JSONString     `json:"settlePx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly"`
		InstType okex.InstrumentType `json:"instType"`
		Delta    okex.JSONString     `json:"delta"`
		Gamma    okex.JSONString     `json:"gamma"`
		Vega     okex.JSONString     `json:"vega"`
		Theta    okex.JSONString     `json:"theta"`
		DeltaBS  okex.JSONString     `json:"deltaBS"`
		GammaBS  okex.JSONString     `json:"gammaBS"`
		VegaBS   okex.JSONString     `json:"vegaBS"`
		ThetaBS  okex.JSONString     `json:"thetaBS"`
		Lever    okex.JSONString     `json:"lever"`
		MarkVol  okex.JSONString     `json:"markVol"`
		BidVol   okex.JSONString     `json:"bidVol"`
		AskVol   okex.JSONString     `json:"askVol"`
		RealVol  okex.JSONString     `json:"realVol"`
		TS       okex.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string          `json:"ccy"`
		Amt          okex.JSONString `json:"amt"`
		DiscountLv   okex.JSONInt64  `json:"discountLv"`
		DiscountInfo []*DiscountInfo `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate okex.JSONInt64 `json:"discountRate"`
		MaxAmt       okex.JSONInt64 `json:"maxAmt"`
		MinAmt       okex.JSONInt64 `json:"minAmt"`
	}
	SystemTime struct {
		TS okex.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okex.InstrumentType       `json:"instType"`
		TotalLoss okex.JSONString           `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string            `json:"ccy,omitempty"`
		Side    okex.OrderSide    `json:"side"`
		OosSide okex.PositionSide `json:"posSide"`
		BkPx    okex.JSONString   `json:"bkPx"`
		Sz      okex.JSONString   `json:"sz"`
		BkLoss  okex.JSONString   `json:"bkLoss"`
		TS      okex.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		MarkPx   okex.JSONString     `json:"markPx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string              `json:"instId"`
		Uly          string              `json:"uly,omitempty"`
		InstType     okex.InstrumentType `json:"instType"`
		Tier         okex.JSONInt64      `json:"tier"`
		MinSz        okex.JSONString     `json:"minSz"`
		MaxSz        okex.JSONString     `json:"maxSz"`
		Mmr          okex.JSONString     `json:"mmr"`
		Imr          okex.JSONString     `json:"imr"`
		OptMgnFactor okex.JSONString     `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan okex.JSONString     `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  okex.JSONString     `json:"baseMaxLoan,omitempty"`
		MaxLever     okex.JSONString     `json:"maxLever"`
		TS           okex.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string          `json:"ccy"`
		Rate  okex.JSONString `json:"rate"`
		Quota okex.JSONString `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string          `json:"level"`
		IrDiscount    okex.JSONString `json:"irDiscount"`
		LoanQuotaCoef int             `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string        `json:"title"`
		State       string        `json:"state"`
		Href        string        `json:"href"`
		ServiceType string        `json:"serviceType"`
		System      string        `json:"system"`
		ScheDesc    string        `json:"scheDesc"`
		Begin       okex.JSONTime `json:"begin"`
		End         okex.JSONTime `json:"end"`
	}
)
