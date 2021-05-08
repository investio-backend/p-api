package model

import (
	"time"

	"gorm.io/gorm"
)

type Fund struct {
	ID               uint16         `json:"-"`
	FundID           string         `json:"fund_id"`
	Code             string         `json:"code"`
	NameEn           string         `json:"name_en"`
	NameTh           string         `json:"name_th"`
	IsPredict        bool           `json:"is_predict"`
	IsFnpick         bool           `json:"-"`
	IsTradable       bool           `json:"is_tradable"`
	IsDividendPayout bool           `json:"is_dividend_payout"`
	FactsheetURL     string         `json:"factsheet_url"`
	InvestStrategyEn string         `json:"invest_strategy_en"`
	InvestStrategyTh string         `json:"invest_strategy_th"`
	ShortDescEn      string         `json:"short_desc_en"`
	ShortDescTh      string         `json:"short_desc_th"`
	InceptionDate    Date           `json:"inception_date"`
	AmcID            uint32         `json:"-"`
	StatID           uint32         `json:"-"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}

// FundSearchResponse - ws
type FundSearchResponse struct {
	FundID string `json:"fund_id"`
	Code   string `json:"code"`
	NameEn string `json:"name_en"`
	NameTh string `json:"name_th"`
}

// type FundAllInfo struct {
// 	FundID           string    `json:"fund_id"`
// 	Code             string    `json:"code"`
// 	NameEn           string    `json:"name_en"`
// 	NameTh           string    `json:"name_th"`
// 	IsPredict        bool      `json:"is_predict"`
// 	IsFnpick         bool      `json:"is_fnpick"`
// 	IsDividendPayout bool      `json:"is_dividend_payout"`
// 	FactsheetURL     string    `json:"factsheet_url"`
// 	InvestStrategyEn string    `json:"invest_strategy_en"`
// 	InvestStrategyTh string    `json:"invest_strategy_th"`
// 	ShortDescEn      string    `json:"short_desc_en"`
// 	ShortDescTh      string    `json:"short_desc_th"`
// 	InceptionDate    time.Time `json:"inception_date"`
// 	AmcCode          string    `json:"amc_code"`
// 	// AmcNameEn        string    `json:"amc_name_en"`
// 	// AmcNameTh        string    `json:"amc_name_th"`
// 	CatID string `json:"cat_id"`
// }

// TableName fund
func (Fund) TableName() string {
	return "fund"
}

// TableName fund
func (FundSearchResponse) TableName() string {
	return "fund"
}

// // TableName fund
// func (FundAllInfo) TableName() string {
// 	return "fund"
// }
