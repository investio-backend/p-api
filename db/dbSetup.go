package db

import (
	"fmt"
	"log"
	"strconv"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"

	"github.com/jinzhu/gorm"
)

var (
	MySQL        *gorm.DB
	InfluxClient influxdb2.Client
	InfluxQuery  api.QueryAPI
)

type DbConfig struct {
	Host     string
	Port     uint64
	User     string
	DbName   string
	Password string
}

func BuildDbConfig(host string, port string, user, pwd, dbName string) *DbConfig {
	portUint, err := strconv.ParseUint(port, 10, 32)
	if err != nil {
		log.Fatalln("Build DB Config: ", err)
	}
	dbConfig := DbConfig{
		Host:     host,
		Port:     portUint,
		User:     user,
		Password: pwd,
		DbName:   dbName,
	}
	return &dbConfig
}

func MySqlURL(dbConfig *DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)
}
