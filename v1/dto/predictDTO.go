package dto

import (
	"github.com/shopspring/decimal"
	"gitlab.com/investio/backend/api/v1/model"
)

type PredictPostReq struct {
	FundCode string          `json:"fund_code"`
	Fiid     uint            `json:"fiid"`
	DataDate model.Date      `json:"data_date" gorm:"type:date;"`
	Prob     decimal.Decimal `json:"buy_prob"`
}
