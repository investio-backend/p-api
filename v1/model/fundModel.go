package model

import (
	"time"

	"gorm.io/gorm"
)

type Fund struct {
	ID               uint32         `json:"id"`
	Code             string         `json:"code"`
	NameEn           string         `json:"name_en"`
	NameTh           string         `json:"name_th"`
	IsPredict        bool           `json:"is_predict"`
	IsFnpick         bool           `json:"is_fnpick"`
	IsDividendPayout bool           `json:"is_dividend_payout"`
	FactsheetURL     string         `json:"factsheet_url"`
	InvestStrategyEn string         `json:"invest_strategy_en"`
	InvestStrategyTh string         `json:"invest_strategy_th"`
	ShortDescEn      string         `json:"short_desc_en"`
	ShortDescTh      string         `json:"short_desc_th"`
	InceptionDate    time.Time      `json:"inception_date"`
	AmcID            uint32         `json:"amc_id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

// FundSearchResponse - ws
type FundSearchResponse struct {
	ID     uint32 `json:"id"`
	Code   string `json:"code"`
	NameEn string `json:"name_en"`
	NameTh string `json:"name_th"`
}

type FundAllInfo struct {
	ID               uint32    `json:"id"`
	Code             string    `json:"code"`
	NameEn           string    `json:"name_en"`
	NameTh           string    `json:"name_th"`
	IsPredict        bool      `json:"is_predict"`
	IsFnpick         bool      `json:"is_fnpick"`
	IsDividendPayout bool      `json:"is_dividend_payout"`
	FactsheetURL     string    `json:"factsheet_url"`
	InvestStrategyEn string    `json:"invest_strategy_en"`
	InvestStrategyTh string    `json:"invest_strategy_th"`
	ShortDescEn      string    `json:"short_desc_en"`
	ShortDescTh      string    `json:"short_desc_th"`
	InceptionDate    time.Time `json:"inception_date"`
}

// TableName fund
func (Fund) TableName() string {
	return "fund"
}

// TableName fund
func (FundSearchResponse) TableName() string {
	return "fund"
}

// TableName fund
func (FundAllInfo) TableName() string {
	return "fund"
}
