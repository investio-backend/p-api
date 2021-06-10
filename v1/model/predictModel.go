package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PredictBuy struct {
	ID        uint            `gorm:"primaryKey" json:"-"`
	FundCode  string          `json:"fund_code"`
	Fiid      uint            `json:"-"`
	DataDate  time.Time       `json:"data_date" gorm:"type:date;"`
	Prob      decimal.Decimal `json:"buy_prob"`
	CreatedAt time.Time       `json:"timestamp"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`
}

func (PredictBuy) TableName() string {
	return "predict_buy"
}

type PredictBuyResponse struct {
	FundID       string          `json:"fund_id"`
	AimcBrdCatID uint            `json:"bcat_id"`
	FundCode     string          `json:"code"`
	RiskID       uint            `json:"risk"`
	DataDate     time.Time       `json:"data_date"`
	Prob         decimal.Decimal `json:"buy_prob"`
}
