package service

import (
	"regexp"

	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/model"
)

type FundService interface {
	GetAllFunds(fund *[]model.FundAllInfo) (err error)
	GetAllCat(aimc_cats *[]model.AimcCat) (err error)
	GetAllBrdCat(cats *[]model.AimcBrdCat) (err error)
	GetAllAmc(amc *[]model.Amc) (err error)
	GetFundInfoByID(fund *model.FundAllInfo, fundID string) error
	SearchFund(query string, limit int) (result []model.FundSearchResponse, err error)
}

type fundService struct {
	thaiRegEx *regexp.Regexp
}

// NewFundService - A constuctor of FundService
func NewFundService(thaiRegEx *regexp.Regexp) FundService {
	return &fundService{
		thaiRegEx: thaiRegEx,
	}
}

func (service *fundService) GetAllFunds(funds *[]model.FundAllInfo) (err error) {
	if err = db.MySQL.Find(&funds).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) GetAllCat(aimc_cats *[]model.AimcCat) (err error) {
	if err = db.MySQL.Find(&aimc_cats).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) GetAllBrdCat(cats *[]model.AimcBrdCat) (err error) {
	if err = db.MySQL.Find(&cats).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) GetAllAmc(amc *[]model.Amc) (err error) {
	if err = db.MySQL.Find(&amc).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) GetFundInfoByID(fund *model.FundAllInfo, fundID string) (err error) {
	// selectQuery := ""
	// query := db.Model(&model.Fund{}).Select(selectQuery).Joins()
	// if err =

	if err = db.MySQL.Where("fund_id = ?", fundID).First(&fund).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) SearchFund(query string, limit int) (result []model.FundSearchResponse, err error) {
	if service.thaiRegEx.MatchString(query) {
		err = service.searchFundByNameTH(&result, query, limit)
	} else {
		err = service.searchFundByFundCode(&result, query, limit)
		resultLen := len(result)
		if resultLen < limit {
			var resultEN []model.FundSearchResponse
			err = service.searchFundByNameEN(&resultEN, query, limit-resultLen)
			result = append(result, resultEN...)
		}
	}
	return
}

func (service *fundService) searchFundByFundCode(funds *[]model.FundSearchResponse, code string, limit int) (err error) {
	if err = db.MySQL.Limit(5).Where("code LIKE ?", "%"+code+"%").Find(&funds).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) searchFundByNameEN(funds *[]model.FundSearchResponse, name string, limit int) (err error) {
	if err = db.MySQL.Limit(limit).Where("name_en LIKE ?", "%"+name+"%").Find(&funds).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) searchFundByNameTH(funds *[]model.FundSearchResponse, name string, limit int) (err error) {
	if err = db.MySQL.Limit(limit).Where("name_th LIKE ?", "%"+name+"%").Find(&funds).Error; err != nil {
		return err
	}
	return nil
}
