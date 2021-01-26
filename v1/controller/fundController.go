package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/dto"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
)

type FundController interface {
	GetFundByID(ctx *gin.Context)
	SearchFund(reqJSON dto.SocketDTO) (response []byte)
}

type fundController struct {
	fundService service.FundService
}

// NewFundController - A constructor of FundController
func NewFundController(service service.FundService) FundController {
	return &fundController{
		fundService: service,
	}
}

func (c *fundController) SearchFund(reqJSON dto.SocketDTO) (response []byte) {
	funds, err := c.fundService.SearchFund(reqJSON.Data)
	if err != nil {
		// panic(err)
		log.Fatal(err)
		errResponse := &dto.SocketDTO{
			Type:  "ERROR",
			Topic: "Database",
			Data:  err.Error(),
		}
		response, err = json.Marshal(errResponse)
		if err != nil {
			log.Fatal("Marshall DB Fail:" + err.Error())
		}
		// response = []byte("Failed: Database " + err.Error())
	} else {
		response, _ = json.Marshal(funds)
		fmt.Println(reqJSON.Topic, reqJSON.Data)
	}
	return
}

func (c *fundController) GetFundByID(ctx *gin.Context) {
	code := ctx.Params.ByName("id")
	var fund model.Fund

	err := c.fundService.GetFundByID(&fund, code)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, fund)
	}
}
