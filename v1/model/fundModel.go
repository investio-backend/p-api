package model

import (
	"time"

	"gorm.io/gorm"
)

type Fund struct {
	ID                uint16         `json:"-"`
	FundID            string         `json:"fund_id"`
	Code              string         `json:"code"`
	NameEn            string         `json:"name_en"`
	NameTh            string         `json:"name_th"`
	IsPredict         bool           `json:"is_predict"`
	IsFnpick          bool           `json:"-"`
	IsTradable        bool           `json:"is_tradable"`
	IsDividendPayout  bool           `json:"is_dividend_payout"`
	FactsheetURL      string         `json:"factsheet_url"`
	ProspectusURL     string         `json:"prospectus_url"`
	HalfyearReportURL string         `json:"halfyear_report_url"`
	AnnualReportURL   string         `json:"annual_report_url"`
	InvestStrategyEn  string         `json:"invest_strategy_en"`
	InvestStrategyTh  string         `json:"invest_strategy_th"`
	ShortDescEn       string         `json:"short_desc_en"`
	ShortDescTh       string         `json:"short_desc_th"`
	InceptionDate     time.Time      `json:"inception_date"`
	AimcBrdCatID      uint           `json:"bcat_id"`
	AmcID             uint32         `json:"-"`
	StatID            uint32         `json:"-"`
	CreatedAt         time.Time      `json:"-"`
	UpdatedAt         time.Time      `json:"-"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

// FundSearchResponse - ws
type FundSearchResponse struct {
	FundID       string `json:"fund_id"`
	Code         string `json:"code"`
	NameEn       string `json:"name_en"`
	NameTh       string `json:"name_th"`
	AimcBrdCatID uint   `json:"bcat_id"`
}

type FundInfoResponse struct {
	FundID            string `json:"fund_id"`
	Code              string `json:"code"`
	NameEn            string `json:"name_en"`
	NameTh            string `json:"name_th"`
	IsPredict         bool   `json:"is_predict"`
	IsDividendPayout  bool   `json:"is_dividend_payout"`
	IsTradable        bool   `json:"is_tradable"`
	FactsheetURL      string `json:"factsheet_url"`
	ProspectusURL     string `json:"prospectus_url"`
	HalfyearReportURL string `json:"halfyear_report_url"`
	AnnualReportURL   string `json:"annual_report_url"`
	InvestStrategyEn  string `json:"invest_strategy_en"`
	InvestStrategyTh  string `json:"invest_strategy_th"`
	ShortDescEn       string `json:"short_desc_en"`
	ShortDescTh       string `json:"short_desc_th"`
	InceptionDate     Date   `json:"inception_date"`
	AimcBrdCatID      uint32 `json:"brd_cat_id"`
	AmcCode           string `json:"amc_code"`
	AmcNameEn         string `json:"amc_name_en"`
	AmcNameTh         string `json:"amc_name_th"`
	CatID             string `json:"cat_id"`
	CatNameEn         string `json:"cat_name_en"`
	CatNameTh         string `json:"cat_name_th"`
	RiskID            uint32 `json:"risk_id"`
	SpectrumEn        string `json:"risk_en"`
	SpectrumTh        string `json:"risk_th"`
	RiskExID          uint32 `json:"risk_exchange_id"`
	TaxTypeID         uint32 `json:"tax_type_id"`
}

// TableName fund
func (Fund) TableName() string {
	return "fund"
}

// TableName fund
func (FundSearchResponse) TableName() string {
	return "fund"
}

type FundScopeResponse struct {
	ID         uint16 `json:"fiid"`
	FundID     string `json:"fund_id"`
	Code       string `json:"code"`
	NameEn     string `json:"name_en"`
	NameTh     string `json:"name_th"`
	IsPredict  bool   `json:"-"`
	IsTradable bool   `json:"-"`
}

// TableName fund
func (FundScopeResponse) TableName() string {
	return "fund"
}
