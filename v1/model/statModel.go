package model

import (
	"time"

	// "github.com/shopspring/decimal"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Stat struct {
	ID              uint16          `json:"id"`
	DataDate        time.Time       `json:"data_date"`
	TotalReturn_1y  decimal.Decimal `json:"total_return_1y" sql:"type:decimal(7,4);"`
	TotalReturnP_1y uint16          `json:"total_return_p_1y"`
	NetAssets       decimal.Decimal `json:"net_asset"`
	Std_1y          decimal.Decimal `json:"std_1y"`
	StdP_1y         uint16          `json:"std_p_1y"`
	StdAvg_1y       decimal.Decimal `json:"std_avg_1y"`
	UnitChange_1y   decimal.Decimal `json:"unit_change_1y"`
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

type Stat_1Y struct {
	DataDate        time.Time       `json:"data_date"`
	TotalReturn_1y  decimal.Decimal `json:"total_return_1y" sql:"type:decimal(7,4);"`
	TotalReturnP_1y uint16          `json:"total_return_p_1y" gorm:"column:total_return_p_1y"`
	NetAssets       decimal.Decimal `json:"net_assets"`
	Std_1y          decimal.Decimal `json:"std_1y"`
	StdP_1y         uint16          `json:"std_p_1y"`
	StdAvg_1y       decimal.Decimal `json:"std_avg_1y"`
	UnitChange_1y   decimal.Decimal `json:"unit_change_1y"`
	FundID          string          `json:"fund_id"`
	Code            string          `json:"code"`
	NameEn          string          `json:"name_en"`
	NameTh          string          `json:"name_th"`
	CatNameEn       string          `json:"cat_name_en"`
}
