package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Stat struct {
	ID                uint32          `json:"id"`
	DataDate          Date            `json:"data_date"`
	TotalReturn_5y    decimal.Decimal `json:"total_return_5y"`
	TotalReturnP_5y   uint16          `json:"total_return_p_5y" gorm:"column:total_return_p_5y"`
	TotalReturnAvg_5y decimal.Decimal `json:"total_return_avg_5y"`
	TotalReturn_3y    decimal.Decimal `json:"total_return_3y"`
	TotalReturnP_3y   uint16          `json:"total_return_p_3y" gorm:"column:total_return_p_3y"`
	TotalReturnAvg_3y decimal.Decimal `json:"total_return_avg_3y"`
	TotalReturn_1y    decimal.Decimal `json:"total_return_1y" sql:"type:decimal(7,4);"`
	TotalReturnP_1y   uint16          `json:"total_return_p_1y" gorm:"column:total_return_p_1y"`
	TotalReturnAvg_1y decimal.Decimal `json:"total_return_avg_1y"`
	TotalReturn_6m    decimal.Decimal `json:"total_return_6m"`
	TotalReturnP_6m   uint16          `json:"total_return_p_6m" gorm:"column:total_return_p_6m"`
	TotalReturnAvg_6m decimal.Decimal `json:"total_return_avg_6m"`
	TotalReturn_3m    decimal.Decimal `json:"total_return_3m"`
	TotalReturnP_3m   uint16          `json:"total_return_p_3m" gorm:"column:total_return_p_3m"`
	TotalReturnAvg_3m decimal.Decimal `json:"total_return_avg_3m"`
	TotalReturn_1m    decimal.Decimal `json:"total_return_1m"`
	TotalReturn_1w    decimal.Decimal `json:"total_return_1w"`
	NetAssets         decimal.Decimal `json:"net_asset"`
	Std_1y            decimal.Decimal `json:"std_1y"`
	StdP_1y           uint16          `json:"std_p_1y" gorm:"column:std_p_1y"`
	StdAvg_1y         decimal.Decimal `json:"std_avg_1y"`
	UnitChange_1y     decimal.Decimal `json:"unit_change_1y"`
	// UnitChange_6m     decimal.Decimal
	// FundID    uint16
	// Fund      FundStatResponse
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName stat
func (Stat) TableName() string {
	return "stat"
}

type StatFundResponse struct {
	FundID            string          `json:"fund_id"`
	Code              string          `json:"fund_code"`
	TotalReturn_5y    decimal.Decimal `json:"total_return_5y"`
	TotalReturnP_5y   uint16          `json:"total_return_p_5y" gorm:"column:total_return_p_5y"`
	TotalReturnAvg_5y decimal.Decimal `json:"total_return_avg_5y"`
	TotalReturn_3y    decimal.Decimal `json:"total_return_3y"`
	TotalReturnP_3y   uint16          `json:"total_return_p_3y" gorm:"column:total_return_p_3y"`
	TotalReturnAvg_3y decimal.Decimal `json:"total_return_avg_3y"`
	TotalReturn_1y    decimal.Decimal `json:"total_return_1y"`
	TotalReturnP_1y   uint16          `json:"total_return_p_1y" gorm:"column:total_return_p_1y"`
	TotalReturnAvg_1y decimal.Decimal `json:"total_return_avg_1y"`
	TotalReturn_6m    decimal.Decimal `json:"total_return_6m"`
	TotalReturnP_6m   uint16          `json:"total_return_p_6m" gorm:"column:total_return_p_6m"`
	TotalReturnAvg_6m decimal.Decimal `json:"total_return_avg_6m"`
	TotalReturn_3m    decimal.Decimal `json:"total_return_3m"`
	TotalReturnP_3m   uint16          `json:"total_return_p_3m" gorm:"column:total_return_p_3m"`
	TotalReturnAvg_3m decimal.Decimal `json:"total_return_avg_3m"`
	TotalReturn_1m    decimal.Decimal `json:"total_return_1m"`
	TotalReturn_1w    decimal.Decimal `json:"total_return_1w"`
	NetAssets         decimal.Decimal `json:"net_asset"`
	Std_1y            decimal.Decimal `json:"std_1y"`
	StdP_1y           uint16          `json:"std_p_1y" gorm:"column:std_p_1y"`
	StdAvg_1y         decimal.Decimal `json:"std_avg_1y"`
	DataDate          Date            `json:"data_date"`
}

type StatTopResponse struct {
	DataDate          Date            `json:"data_date"`
	TotalReturn_5y    decimal.Decimal `json:"total_return_5y"`
	TotalReturnP_5y   uint16          `json:"total_return_p_5y" gorm:"column:total_return_p_5y"`
	TotalReturnAvg_5y decimal.Decimal `json:"total_return_avg_5y"`
	TotalReturn_3y    decimal.Decimal `json:"total_return_3y"`
	TotalReturnP_3y   uint16          `json:"total_return_p_3y" gorm:"column:total_return_p_3y"`
	TotalReturnAvg_3y decimal.Decimal `json:"total_return_avg_3y"`
	TotalReturn_1y    decimal.Decimal `json:"total_return_1y"`
	TotalReturnP_1y   uint16          `json:"total_return_p_1y"  gorm:"column:total_return_p_1y"`
	TotalReturnAvg_1y decimal.Decimal `json:"total_return_avg_1y"`
	TotalReturn_6m    decimal.Decimal `json:"total_return_6m"`
	TotalReturnP_6m   uint16          `json:"total_return_p_6m" gorm:"column:total_return_p_6m"`
	TotalReturnAvg_6m decimal.Decimal `json:"total_return_avg_6m"`
	TotalReturn_3m    decimal.Decimal `json:"total_return_3m"`
	TotalReturnP_3m   uint16          `json:"total_return_p_3m" gorm:"column:total_return_p_3m"`
	TotalReturnAvg_3m decimal.Decimal `json:"total_return_avg_3m"`
	NetAssets         decimal.Decimal `json:"net_assets"`
	FundID            string          `json:"fund_id"`
	Code              string          `json:"code"`
	NameEn            string          `json:"name_en"`
	NameTh            string          `json:"name_th"`
	CatNameEn         string          `json:"cat_name_en"`
	CatNameTh         string          `json:"cat_name_th"`
	AmcCode           string          `json:"amc_code"`
	AmcNameEn         string          `json:"amc_name_en"`
	AmcNameTh         string          `json:"amc_name_th"`
}

type Stat_1Y struct {
	DataDate          Date            `json:"data_date"`
	TotalReturn_1y    decimal.Decimal `json:"total_return_1y"`
	TotalReturnP_1y   uint16          `json:"total_return_p_1y"  gorm:"column:total_return_p_1y"`
	TotalReturnAvg_1y decimal.Decimal `json:"total_return_avg_1y"`

	NetAssets decimal.Decimal `json:"net_assets"`
	FundID    string          `json:"fund_id"`
	Code      string          `json:"code"`
	NameEn    string          `json:"name_en"`
	NameTh    string          `json:"name_th"`
	CatNameEn string          `json:"cat_name_en"`
	CatNameTh string          `json:"cat_name_th"`
	AmcCode   string          `json:"amc_code"`
	AmcNameEn string          `json:"amc_name_en"`
	AmcNameTh string          `json:"amc_name_th"`
	// Std_1y            decimal.Decimal `json:"std_1y"`
	// StdP_1y           uint16          `json:"std_p_1y"`
	// StdAvg_1y         decimal.Decimal `json:"std_avg_1y"`
	// UnitChange_1y decimal.Decimal `json:"unit_change_1y"`
}

type Stat_6M struct {
	DataDate          Date            `json:"data_date"`
	TotalReturn_6m    decimal.Decimal `json:"total_return_6m"`
	TotalReturnP_6m   uint16          `json:"total_return_p_6m" gorm:"column:total_return_p_6m"`
	TotalReturnAvg_6m decimal.Decimal `json:"total_return_avg_6m"`
	NetAssets         decimal.Decimal `json:"net_assets"`
	FundID            string          `json:"fund_id"`
	Code              string          `json:"code"`
	NameEn            string          `json:"name_en"`
	NameTh            string          `json:"name_th"`
	CatNameEn         string          `json:"cat_name_en"`
	CatNameTh         string          `json:"cat_name_th"`
	AmcCode           string          `json:"amc_code"`
	AmcNameEn         string          `json:"amc_name_en"`
	AmcNameTh         string          `json:"amc_name_th"`
	// UnitChange_6m     decimal.Decimal `json:"unit_change_6m"`
}

type Stat_3M struct {
	DataDate Date `json:"data_date"`
}
