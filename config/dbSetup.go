package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Host string
	Port int
	User string
	DbName string
	Password string
}


func BuildDbConfig() *DbConfig {
	dbConfig := DbConfig {
		Host: "192.168.50.121",
		Port: 6006,
		User: "pyFund",
		Password: "L5CJMVvs84",
		DbName: "investio",
	}
	return &dbConfig
}

func DbURL(dbConfig *DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)
}
