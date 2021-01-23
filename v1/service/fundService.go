package service

import (
	"fmt"

	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/model"
	"gopkg.in/olahol/melody.v1"
)

type FundService interface {
	GetAllFunds(fund *[]model.Fund) (err error)
	GetFundByID(fund *model.Fund, fundID string) error
	HandleWsConnect(s *melody.Session)
	SearchFund(query string) (result []model.FundSearchResponse, err error)
}

type fundService struct {
}

// NewFundService - A constuctor of FundService
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

func (service *fundService) HandleWsConnect(s *melody.Session) {
	fmt.Println("Connected: " + s.Request.Host)
}

func (service *fundService) SearchFund(query string) (result []model.FundSearchResponse, err error) {
	err = service.searchFundByFundCode(&result, query)
	return
}

func (service *fundService) searchFundByFundCode(funds *[]model.FundSearchResponse, code string) (err error) {
	if err = db.MySQL.Limit(5).Where("code LIKE ?", "%"+code+"%").Find(funds).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) searchFundByNameEn(funds *[]model.FundSearchResponse, name string) (err error) {
	if err = db.MySQL.Where("name_en LIKE ?", "%"+name+"%").Find(&funds).Error; err != nil {
		return err
	}
	return nil
}
