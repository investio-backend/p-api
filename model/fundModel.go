package model

import "time"

type Fund struct {
	Id               int64     `json:"id"`
	Code             string    `json:"code"`
	NameEn           string    `json:"name_en"`
	NameTh           string    `json:"name_th"`
	IsPredict        bool      `json:"is_predict"`
	IsFnPick         bool      `json:"is_fnpick"`
	IsDividendPayout bool      `json:"is_dividend_payout"`
	FnRef            string    `json:"fn_ref"`
	AimcFundRef      string    `json:"aime_fund_ref"`
	FactsheetURL     string    `json:"factsheet_url"`
	InvestStrategyEn string    `json:"invest_strategy_en"`
	InvestStrategyTh string    `json:"invest_strategy_th"`
	ShortDescEn      string    `json:"short_desc_en"`
	ShortDescTh      string    `json:"short_desc_th"`
	InceptionDate    time.Time `json:"inception_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// type FundBriefInfo struct {
// 	Id int64 `json:"id"`
// 	Code string `json:"code"`
// 	NameEn           string    `json:"name_en"`
// 	NameTh           string    `json:"name_th"`
// }

func (b *Fund) TableName() string {
	return "fund"
}
