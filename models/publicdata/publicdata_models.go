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
		CtVal     string               `json:"ctVal,omitempty"`
		CtMult    string               `json:"ctMult,omitempty"`
		Stk       string               `json:"stk,omitempty"`
		TickSz    string               `json:"tickSz,omitempty"`
		LotSz     string               `json:"lotSz,omitempty"`
		MinSz     string               `json:"minSz,omitempty"`
		Lever     string               `json:"lever"`
		OptType   okex.OptionType      `json:"optType,omitempty"`
		ListTime  okex.JSONTime        `json:"listTime"`
		ExpTime   okex.JSONTime        `json:"expTime,omitempty"`
		CtType    okex.ContractType    `json:"ctType,omitempty"`
		Alias     okex.AliasType       `json:"alias,omitempty"`
		State     okex.InstrumentState `json:"state"`

		MaxLmSz      string `json:"maxLmtSz,omitempty"`
		MaxMkSz      string `json:"maxMktSz,omitempty"`
		MaxTwapSz    string `json:"maxTwapSz,omitempty"`
		MaxIcebergSz string `json:"maxIcebergSz,omitempty"`
		MaxTriggerSz string `json:"maxTriggerSz,omitempty"`
		MaxStopSz    string `json:"maxStopSz,omitempty"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okex.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId"`
		Px     string                    `json:"px"`
		Type   okex.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string              `json:"instId"`
		Oi       string              `json:"oi"`
		OiCcy    string              `json:"oiCcy"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string              `json:"instId"`
		InstType        okex.InstrumentType `json:"instType"`
		FundingRate     string              `json:"fundingRate"`
		NextFundingRate string              `json:"NextFundingRate"`
		FundingTime     okex.JSONTime       `json:"fundingTime"`
		NextFundingTime okex.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		BuyLmt   string              `json:"buyLmt"`
		SellLmt  string              `json:"sellLmt"`
		TS       okex.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		SettlePx string              `json:"settlePx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly"`
		InstType okex.InstrumentType `json:"instType"`
		Delta    string              `json:"delta"`
		Gamma    string              `json:"gamma"`
		Vega     string              `json:"vega"`
		Theta    string              `json:"theta"`
		DeltaBS  string              `json:"deltaBS"`
		GammaBS  string              `json:"gammaBS"`
		VegaBS   string              `json:"vegaBS"`
		ThetaBS  string              `json:"thetaBS"`
		Lever    string              `json:"lever"`
		MarkVol  string              `json:"markVol"`
		BidVol   string              `json:"bidVol"`
		AskVol   string              `json:"askVol"`
		RealVol  string              `json:"realVol"`
		TS       okex.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string          `json:"ccy"`
		Amt          string          `json:"amt"`
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
		TotalLoss string                    `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string            `json:"ccy,omitempty"`
		Side    okex.OrderSide    `json:"side"`
		OosSide okex.PositionSide `json:"posSide"`
		BkPx    string            `json:"bkPx"`
		Sz      string            `json:"sz"`
		BkLoss  string            `json:"bkLoss"`
		TS      okex.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		MarkPx   string              `json:"markPx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string              `json:"instId"`
		Uly          string              `json:"uly,omitempty"`
		InstType     okex.InstrumentType `json:"instType"`
		Tier         okex.JSONInt64      `json:"tier"`
		MinSz        string              `json:"minSz"`
		MaxSz        string              `json:"maxSz"`
		Mmr          string              `json:"mmr"`
		Imr          string              `json:"imr"`
		OptMgnFactor string              `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan string              `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  string              `json:"baseMaxLoan,omitempty"`
		MaxLever     string              `json:"maxLever"`
		TS           okex.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string `json:"ccy"`
		Rate  string `json:"rate"`
		Quota string `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string `json:"level"`
		IrDiscount    string `json:"irDiscount"`
		LoanQuotaCoef int    `json:"loanQuotaCoef,string"`
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
