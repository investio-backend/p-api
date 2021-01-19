package service

import (
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/investio/backend/api/config"
	"gitlab.com/investio/backend/api/model"
)

func GetAllFunds(fund *[]model.Fund) (err error) {
	if err = config.DB.Find(fund).Error; err != nil {
		return err
	}
	return nil
}

func GetFundByFundCode(fund *model.Fund, code string) (err error) {
	if err = config.DB.Where("code = ?", code).First(fund).Error; err != nil {
		return err
	}
	return nil
}
