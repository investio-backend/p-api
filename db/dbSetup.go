package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"gitlab.com/investio/backend/api/v1/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func SetupDB() (err error) {
	MySQL, err = gorm.Open(
		mysql.Open(
			MySqlURL(BuildDbConfig(
				os.Getenv("MYSQL_HOST"),
				os.Getenv("MYSQL_PORT"),
				os.Getenv("MYSQL_USER"),
				os.Getenv("MYSQL_PWD"),
				os.Getenv("MYSQL_DB"),
			)),
		),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatalln("Database Init error: ", err)
	} else {
		MySQL.AutoMigrate(&model.PredictBuy{})

		InfluxClient = influxdb2.NewClient(
			os.Getenv("INFLUX_HOST"),
			os.Getenv("INFLUX_TOKEN"),
		)
		InfluxQuery = InfluxClient.QueryAPI(os.Getenv("INFLUX_ORG"))
	}
	return
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
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)
}
