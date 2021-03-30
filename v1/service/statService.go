package service

import (
	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/model"
)

type StatService interface {
	FindTopStat1Y(stats *[]model.Stat_1Y, catID string, amcCode string) (err error)
	FindTopStat6M(stats *[]model.Stat_6M, catID string, amcCode string) (err error)
	GetStatByFundID(results *model.StatFundResponse, fundID string) (err error)
}

type statService struct {
}

// NewFundService - A constuctor of FundService
func NewStatService() StatService {
	return &statService{}
}

func (service *statService) FindTopStat1Y(result *[]model.Stat_1Y, catID string, amcCode string) (err error) {
	// TODO: data_date: "" & "total_return_p_1y": 0
	selectQ := "stat.net_assets, stat.data_date, fund.code, fund.fund_id, fund.name_en, fund.name_th, aimc_cat.cat_name_en, aimc_cat.cat_name_th, aimc_cat.cat_name_th, amc.amc_code, amc.amc_name_en, amc.amc_name_th, "

	selectQ += "stat.total_return_1y, stat.total_return_p_1y, stat.total_return_avg_1y"
	orderBy := "total_return_1y desc"

	query := db.MySQL.Model(&model.Stat{}).Limit(50).Order(orderBy).Select(selectQ).Joins("join fund on stat.id = fund.stat_id").Joins("join aimc_cat on fund.aimc_cat_id = aimc_cat.id").Joins("join amc on fund.amc_id = amc.id")

	if catID != "" {
		query = query.Where("cat_id = ?", catID)
	}

	if amcCode != "" {
		query = query.Where("amc_code = ?", amcCode)
	}

	if err = query.Find(&result).Error; err != nil {
		return err
	}
	return nil
}

func (service *statService) FindTopStat6M(result *[]model.Stat_6M, catID string, amcCode string) (err error) {
	// TODO: data_date: "" & "total_return_p_1y": 0
	selectQ := `stat.net_assets, stat.data_date, fund.code, fund.fund_id, fund.name_en, fund.name_th, 
		aimc_cat.cat_name_en, aimc_cat.cat_name_th, amc.amc_code, amc.amc_name_en, amc.amc_name_th, `

	selectQ += "stat.total_return_6m, stat.total_return_p_6m, stat.total_return_avg_6m"
	orderBy := "total_return_6m desc"

	query := db.MySQL.Model(&model.Stat{}).Limit(50).Order(orderBy).Select(selectQ)
	query = query.Joins("join fund on stat.id = fund.stat_id").Joins("join aimc_cat on fund.aimc_cat_id = aimc_cat.id").Joins("join amc on fund.amc_id = amc.id")
	// fmt.Println(catID)
	if catID != "" {
		query = query.Where("cat_id = ?", catID)
	}

	if amcCode != "" {
		query = query.Where("amc_code = ?", amcCode)
	}

	if err = query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func (s *statService) GetStatByFundID(results *model.StatFundResponse, fundID string) (err error) {
	selectStr := `fund.fund_id, fund.code, stat.total_return_1y, stat.total_return_p_1y, stat.total_return_avg_1y, 
		stat.total_return_6m, stat.total_return_p_6m, stat.total_return_avg_6m, stat.total_return_3m, stat.total_return_p_3m, stat.total_return_avg_3m, 
		stat.total_return_1m, stat.total_return_1w, stat.net_assets, stat.std_1y, stat.std_p_1y, stat.std_avg_1y, stat.data_date`

	err = db.MySQL.Model(&model.Fund{}).Select(selectStr).Joins("join stat on fund.stat_id = stat.id").First(&results, "fund.fund_id = ?", fundID).Error
	return
}
