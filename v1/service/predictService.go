package service

import (
	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/model"
)

type PredictService interface {
	WriteResult(result *model.PredictBuy) (err error)
	ReadTopResult(results *[]model.PredictBuyResponse) (err error)
}

type predictService struct {
}

// NewPredictService - A constuctor
func NewPredictService() PredictService {
	return &predictService{}
}

func (s *predictService) WriteResult(result *model.PredictBuy) (err error) {
	err = db.MySQL.Create(result).Error
	return
}

func (s *predictService) ReadTopResult(results *[]model.PredictBuyResponse) (err error) {
	selectStr := "fund.fund_id, fund.aimc_brd_cat_id, fund.risk_id, predict_buy.fund_code, predict_buy.data_date, predict_buy.prob"
	query := db.MySQL.Model(&model.PredictBuy{}).Limit(10).Order("prob desc")
	err = query.Select(selectStr).Joins("join fund on predict_buy.fiid = fund.id").Find(results).Error
	return
}
