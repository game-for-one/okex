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
		CtVal     okex.JSONAmount      `json:"ctVal,omitempty"`
		CtMult    okex.JSONAmount      `json:"ctMult,omitempty"`
		Stk       okex.JSONAmount      `json:"stk,omitempty"`
		TickSz    okex.JSONAmount      `json:"tickSz,omitempty"`
		LotSz     okex.JSONAmount      `json:"lotSz,omitempty"`
		MinSz     okex.JSONAmount      `json:"minSz,omitempty"`
		Lever     okex.JSONAmount      `json:"lever"`
		OptType   okex.OptionType      `json:"optType,omitempty"`
		ListTime  okex.JSONTime        `json:"listTime"`
		ExpTime   okex.JSONTime        `json:"expTime,omitempty"`
		CtType    okex.ContractType    `json:"ctType,omitempty"`
		Alias     okex.AliasType       `json:"alias,omitempty"`
		State     okex.InstrumentState `json:"state"`

		MaxLmSz      okex.JSONAmount `json:"maxLmtSz,omitempty"`
		MaxMkSz      okex.JSONAmount `json:"maxMktSz,omitempty"`
		MaxTwapSz    okex.JSONAmount `json:"maxTwapSz,omitempty"`
		MaxIcebergSz okex.JSONAmount `json:"maxIcebergSz,omitempty"`
		MaxTriggerSz okex.JSONAmount `json:"maxTriggerSz,omitempty"`
		MaxStopSz    okex.JSONAmount `json:"maxStopSz,omitempty"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okex.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId"`
		Px     okex.JSONAmount           `json:"px"`
		Type   okex.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string              `json:"instId"`
		Oi       okex.JSONAmount     `json:"oi"`
		OiCcy    okex.JSONAmount     `json:"oiCcy"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string              `json:"instId"`
		InstType        okex.InstrumentType `json:"instType"`
		FundingRate     okex.JSONAmount     `json:"fundingRate"`
		NextFundingRate okex.JSONAmount     `json:"NextFundingRate"`
		FundingTime     okex.JSONTime       `json:"fundingTime"`
		NextFundingTime okex.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		BuyLmt   okex.JSONAmount     `json:"buyLmt"`
		SellLmt  okex.JSONAmount     `json:"sellLmt"`
		TS       okex.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		SettlePx okex.JSONAmount     `json:"settlePx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly"`
		InstType okex.InstrumentType `json:"instType"`
		Delta    okex.JSONAmount     `json:"delta"`
		Gamma    okex.JSONAmount     `json:"gamma"`
		Vega     okex.JSONAmount     `json:"vega"`
		Theta    okex.JSONAmount     `json:"theta"`
		DeltaBS  okex.JSONAmount     `json:"deltaBS"`
		GammaBS  okex.JSONAmount     `json:"gammaBS"`
		VegaBS   okex.JSONAmount     `json:"vegaBS"`
		ThetaBS  okex.JSONAmount     `json:"thetaBS"`
		Lever    okex.JSONAmount     `json:"lever"`
		MarkVol  okex.JSONAmount     `json:"markVol"`
		BidVol   okex.JSONAmount     `json:"bidVol"`
		AskVol   okex.JSONAmount     `json:"askVol"`
		RealVol  okex.JSONAmount     `json:"realVol"`
		TS       okex.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string          `json:"ccy"`
		Amt          okex.JSONAmount `json:"amt"`
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
		TotalLoss okex.JSONAmount           `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string            `json:"ccy,omitempty"`
		Side    okex.OrderSide    `json:"side"`
		OosSide okex.PositionSide `json:"posSide"`
		BkPx    okex.JSONAmount   `json:"bkPx"`
		Sz      okex.JSONAmount   `json:"sz"`
		BkLoss  okex.JSONAmount   `json:"bkLoss"`
		TS      okex.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		MarkPx   okex.JSONAmount     `json:"markPx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string              `json:"instId"`
		Uly          string              `json:"uly,omitempty"`
		InstType     okex.InstrumentType `json:"instType"`
		Tier         okex.JSONInt64      `json:"tier"`
		MinSz        okex.JSONAmount     `json:"minSz"`
		MaxSz        okex.JSONAmount     `json:"maxSz"`
		Mmr          okex.JSONAmount     `json:"mmr"`
		Imr          okex.JSONAmount     `json:"imr"`
		OptMgnFactor okex.JSONAmount     `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan okex.JSONAmount     `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  okex.JSONAmount     `json:"baseMaxLoan,omitempty"`
		MaxLever     okex.JSONAmount     `json:"maxLever"`
		TS           okex.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string          `json:"ccy"`
		Rate  okex.JSONAmount `json:"rate"`
		Quota okex.JSONAmount `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string          `json:"level"`
		IrDiscount    okex.JSONAmount `json:"irDiscount"`
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
