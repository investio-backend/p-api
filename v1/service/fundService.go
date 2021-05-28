package service

import (
	"regexp"

	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/model"
)

type FundService interface {
	// GetAllFunds(fund *[]model.Fund) (err error)
	GetAllCat(aimc_cats *[]model.AimcBrdCat) (err error)
	// GetAllBrdCat(cats *[]model.AimcBrdCat) (err error)
	GetAllAmc(amc *[]model.Amc) (err error)
	GetFundInfoByID(results *model.FundInfoResponse, fundID string) error
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

// func (service *fundService) GetAllFunds(funds *[]model.Fund) (err error) {
// 	if err = db.MySQL.Find(&funds).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (service *fundService) GetAllCat(aimc_cats *[]model.AimcCat) (err error) {
// 	if err = db.MySQL.Find(&aimc_cats).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (service *fundService) GetAllBrdCat(cats *[]model.AimcBrdCat) (err error) {
// 	if err = db.MySQL.Find(&cats).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func (service *fundService) GetAllCat(aimc_cats *[]model.AimcBrdCat) (err error) {
	err = db.MySQL.Preload("Cats").Find(&aimc_cats).Error
	return
}

func (service *fundService) GetAllAmc(amc *[]model.Amc) (err error) {
	if err = db.MySQL.Find(&amc).Error; err != nil {
		return err
	}
	return nil
}

func (service *fundService) GetFundInfoByID(results *model.FundInfoResponse, fundID string) (err error) {
	selectStr := `fund.fund_id, fund.code, fund.name_en, fund.name_th, fund.is_predict, fund.is_dividend_payout, fund.is_tradable,
					fund.factsheet_url, fund.prospectus_url, fund.halfyear_report_url, fund.annual_report_url, 
					fund.invest_strategy_en, fund.invest_strategy_th, fund.short_desc_en, fund.short_desc_th, fund.aimc_brd_cat_id,
					fund.inception_date, aimc_cat.cat_id, aimc_cat.cat_name_en, aimc_cat.cat_name_th,
					amc.amc_code, amc.amc_name_en, amc.amc_name_th, fund.risk_id, fund.risk_ex_id, fund.tax_type_id`
	// selectQuery := ""
	// query := db.Model(&model.Fund{}).Select(selectQuery).Joins()
	// if err =

	// if err = db.MySQL.Where("fund_id = ?", fundID).First(&fund).Error; err != nil {
	// 	return err
	// }
	// return nil

	// selectQ := `fund`

	err = db.MySQL.Model(&model.Fund{}).Select(selectStr).Joins("join aimc_cat on fund.aimc_cat_id = aimc_cat.id").Joins("join amc on fund.amc_id = amc.id").First(&results, "fund.fund_id = ?", fundID).Error
	return
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
