package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
		config.DbURL(config.BuildDbConfig()),
	)

	if err != nil {
		fmt.Println("Database Status: ", err)
	}

	config.DB.AutoMigrate(&model.Fund{})
}

func main() {
	SetupDB()
	defer config.DB.Close()

	server := gin.Default()

	routeV0 := server.Group("/v0")
	{
		routeV0.GET("fund/:code", fundController.GetFundByCode)
	}

	server.Run()
}
