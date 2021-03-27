package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/dto"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
	"gorm.io/gorm"
)

// FundController manages fund
type FundController interface {
	GetFundByID(ctx *gin.Context)
	ListCat(ctx *gin.Context)
	ListAmc(ctx *gin.Context)
	// GetAllFund(ctx *gin.Context)
	SearchFund(ctx *gin.Context)
	SocketSearchFund(reqJSON dto.SocketDTO) (response []byte)
	GetTopReturn(ctx *gin.Context)
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

// TODO: Remove
func (c *fundController) SocketSearchFund(reqJSON dto.SocketDTO) (response []byte) {
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

func (c *fundController) SearchFund(ctx *gin.Context) {
	query := ctx.Params.ByName("fundQuery")
	LIMIT := 5

	if funds, err := c.fundService.SearchFund(query, LIMIT); err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
	} else {
		ctx.JSON(http.StatusOK, funds)
	}
}

func (c *fundController) GetFundByID(ctx *gin.Context) {
	code := ctx.Params.ByName("id")
	var fund model.FundAllInfo

	err := c.fundService.GetFundInfoByID(&fund, code)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}

	} else {
		// fmt.Println(fund)
		ctx.JSON(http.StatusOK, fund)
	}
}

func (c *fundController) ListCat(ctx *gin.Context) {
	var catList []model.AimcCat
	if err := c.fundService.GetAllCat(&catList); err == nil {
		ctx.JSON(http.StatusOK, catList)
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}

func (c *fundController) ListAmc(ctx *gin.Context) {
	var amcList []model.Amc
	if err := c.fundService.GetAllAmc(&amcList); err == nil {
		ctx.JSON(http.StatusOK, amcList)
	}
}

func (c *fundController) GetTopReturn(ctx *gin.Context) {

	var queryStr dto.QueryStrStat

	if ctx.ShouldBind(&queryStr) != nil {
		queryStr = dto.QueryStrStat{
			Amc:   "",
			Cat:   "",
			Range: "1y",
		}
	}
	if strings.ToLower(queryStr.Range) == "1y" {
		var statRes []model.Stat_1Y
		if err := c.fundService.FindTopStat1Y(&statRes, queryStr.Cat, queryStr.Amc); err != nil {
			// fmt.Println("Return Err: ", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			// fmt.Println("Res: ", statRes)
			ctx.JSON(http.StatusOK, statRes)
		}
	} else if strings.ToLower(queryStr.Range) == "6m" {
		var statRes []model.Stat_6M
		if err := c.fundService.FindTopStat6M(&statRes, queryStr.Cat, queryStr.Amc); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			ctx.JSON(http.StatusOK, statRes)
		}
	} else {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
	}
}
