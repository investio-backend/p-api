package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/dto"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
	"gorm.io/gorm"
)

// FundController manages fund
type FundController interface {
	GetFundByID(ctx *gin.Context)
	// GetAllFund(ctx *gin.Context)
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
	LIMIT := 5
	funds, err := c.fundService.SearchFund(reqJSON.Data, LIMIT)
	// fmt.Println("Search: ", funds)
	if err != nil {
		// panic(err)
		log.Fatal(err)
		errResponse := &dto.SocketDTO{
			Type:  "ERROR",
			Topic: "database",
			Data:  err.Error(),
		}
		response, err = json.Marshal(errResponse)
		if err != nil {
			log.Fatal("Marshall DB Fail:" + err.Error())
		}
		// response = []byte("Failed: Database " + err.Error())
	} else if len(funds) > 0 {
		// fundsData, _ := json.Marshal(funds)
		dataR := &dto.SocketArrayDTO{
			Type:  "FUNDRES",
			Topic: "search",
			Data:  funds,
		}
		response, _ = json.Marshal(dataR)
		// fmt.Println(reqJSON.Topic, reqJSON.Data)
	}
	return
}

func (c *fundController) GetFundByID(ctx *gin.Context) {
	code := ctx.Params.ByName("id")
	var fund model.FundAllInfo

	err := c.fundService.GetFundInfoByID(&fund, code)

	fmt.Println(fund)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatus(http.StatusBadGateway)
		}

	} else {
		ctx.JSON(http.StatusOK, fund)
	}
}
