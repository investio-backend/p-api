package service

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/model"
	"gopkg.in/olahol/melody.v1"
)

type FundService interface {
	GetAllFunds(fund *[]model.Fund) (err error)
	SearchFundByFundCode(fund *[]model.FundSearch, code string) (err error)
	GetFundByID(fund *model.Fund, fundID string) error
	HandleConnect(s *melody.Session)
}

type fundService struct {
	cachedFund model.Fund
}

func NewFundService() FundService {
	return &fundService{}
}

func (service *fundService) GetAllFunds(fund *[]model.Fund) (err error) {
	if err = db.MySQL.Find(fund).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) GetFundByID(fund *model.Fund, fundID string) error {
	if err := db.MySQL.Where("id = ?", fundID).First(&fund).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) SearchFundByFundCode(funds *[]model.FundSearch, code string) (err error) {
	if err = db.MySQL.Where("code LIKE ?", "%"+code+"%").Find(funds).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) SearchFundByNameEn(funds *[]model.FundSearch, query string) (err error) {
	if err = db.MySQL.Where("name_en LIKE ?", "%"+query+"%").Find(&funds).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) HandleConnect(s *melody.Session) {
	fmt.Println("Connected: " + s.Request.Host)
}
