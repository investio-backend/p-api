package service

import (
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/model"
)

type FundService interface {
	GetAllFunds(fund []*model.Fund) (err error)
	GetFundByFundCode(fund *model.Fund, code string) (err error)
}

type fundService struct {
	cachedFund model.Fund
}

func New() FundService {
	return &fundService{}
}

func (service *fundService) GetAllFunds(fund []*model.Fund) (err error) {
	if err = db.MySQL.Find(fund).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) GetFundByFundCode(fund *model.Fund, code string) (err error) {
	if err = db.MySQL.Where("code LIKE ?", "%"+code+"%").First(fund).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) GetFundByNameEn(funds []*model.Fund, query string) (err error) {
	if err = db.MySQL.Where("name_eb LIKE ?", "%"+query+"%").Find(&funds).Error; err != nil {
		return err
	}
	return nil
}
