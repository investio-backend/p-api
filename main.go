package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gitlab.com/investio/backend/api/config"
	"gitlab.com/investio/backend/api/controller"
	"gitlab.com/investio/backend/api/model"
	"gitlab.com/investio/backend/api/service"
)

var (
	err         error
	fundService service.FundService = service.New()

	fundController controller.FundController = controller.New(fundService)
)

func SetupDB() {
	config.DB, err = gorm.Open(
		"mysql",
		config.MySqlURL(config.BuildDbConfig(
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

	config.DB.AutoMigrate(&model.Fund{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SetupDB()
	defer config.DB.Close()

	server := gin.Default()

	routeV0 := server.Group("/v0")
	{
		routeV0.GET("fund/:code", fundController.GetFundByCode)
	}

	server.Run(":" + os.Getenv("API_PORT"))
}
