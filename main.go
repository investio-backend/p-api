package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gitlab.com/investio/backend/api/controller"
	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/model"
	"gitlab.com/investio/backend/api/service"
)

var (
	err error

	fundService service.FundService = service.New()
	navService  service.NavService  = service.NewNavService()

	fundController controller.FundController = controller.New(fundService)
	navController  controller.NavController  = controller.NewNavController(navService)
)

func setupDB() {
	db.MySQL, err = gorm.Open(
		"mysql",
		db.MySqlURL(db.BuildDbConfig(
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PWD"),
			os.Getenv("MYSQL_DB"),
		)),
	)

	if err != nil {
		fmt.Println("Database Status: ", err)
	}

	db.MySQL.AutoMigrate(&model.Fund{})

	db.InfluxClient = influxdb2.NewClient(
		"http://"+os.Getenv("INFLUX_HOST")+":"+os.Getenv("INFLUX_PORT"),
		os.Getenv("INFLUX_TOKEN"),
	)
	db.InfluxQuery = db.InfluxClient.QueryAPI("Investio")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	setupDB()
	defer db.MySQL.Close()
	defer db.InfluxClient.Close()

	fmt.Printf("Type: %T", db.MySQL)

	server := gin.Default()

	routeV0 := server.Group("/v0")
	{
		routeV0.GET("funds/:code", fundController.GetFundByCode)
		routeV0.GET("navs/:code", navController.GetPastNavSeries)
	}

	server.Run(":" + os.Getenv("API_PORT"))
}
