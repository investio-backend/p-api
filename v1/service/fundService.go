package service

import (
	"fmt"
	"regexp"

	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/model"
)

type FundService interface {
	GetAllFunds(fund *[]model.FundAllInfo) (err error)
	GetAllCat(aimc_cats *[]model.AimcCat) (err error)
	GetAllBrdCat(cats *[]model.AimcBrdCat) (err error)
	GetFundInfoByID(fund *model.FundAllInfo, fundID string) error
	SearchFund(query string, limit int) (result []model.FundSearchResponse, err error)
	FindTopReturn(stats *[]model.Stat_1Y, catID string, duration string) (err error)
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

func (service *fundService) GetFundInfoByID(fund *model.FundAllInfo, fundID string) (err error) {
	if err = db.MySQL.Where("fund_id = ?", fundID).First(&fund).Error; err != nil {
		return err
	}
	return nil
	// if err := db.MySQL.Joins("JOIN")
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

func (service *fundService) FindTopReturn(result *[]model.Stat_1Y, catID string, duration string) (err error) {
	// var fund model.Fund
	// if err = db.MySQL.Model(&fund).Association("Stat").Error; err != nil {
	// 	return err
	// }
	// db.MySQL.Model(&fund).Limit(50).Order("total_return_1y desc").Association("Stat").Find(&stats)
	// return nil

	// if err = db.MySQL.Limit(50).Preload("Fund").Order("total_return_1y desc").Find(&stats).Error; err != nil {
	// 	return err
	// }
	// return nil

	selectQ := "stat.std_1y, stat.net_assets, fund.code, fund.name_en, fund.name_th, aimc_cat.cat_name_en, "

	selectQ += "stat.total_return_1y, stat.total_return_p_1y"

	query := db.MySQL.Model(&model.Stat{}).Limit(50).Order("total_return_1y desc").Select(selectQ).Joins("join fund on stat.id = fund.stat_id").Joins("join aimc_cat on fund.aimc_cat_id = aimc_cat.id")
	fmt.Println(catID)
	if catID == "" {
		if err = query.Scan(&result).Error; err != nil {
			return err
		}
	} else {
		if err = query.Where("cat_id = ?", catID).Scan(&result).Error; err != nil {
			return err
		}
	}
	return nil
}
