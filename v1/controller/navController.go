package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/dto"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
)

// NavController manages NAV
type NavController interface {
	GetPastNavSeriesByFundCode(ctx *gin.Context)
	GetLatestNav(ctx *gin.Context)
}

type navController struct {
	service service.NavService
}

// NewNavController - A constructor of NavController
func NewNavController(service service.NavService) NavController {
	return &navController{
		service: service,
	}
}

func (c *navController) GetPastNavSeriesByFundCode(ctx *gin.Context) {
	var (
		pastNav   []model.NavDate
		reqByCode pastNavByFundCode
		err       error
	)

	if err = ctx.ShouldBind(&reqByCode); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(reqByCode.FundCode)
	if reqByCode.FundCode != "" {
		err = c.service.GetPastNavByFundCode(&pastNav, reqByCode.FundCode, reqByCode.Range)
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	if pastNav == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	response := model.NavSeries{
		FundCode: reqByCode.FundCode, Navs: pastNav,
	}
	ctx.JSON(http.StatusOK, response)

}

func (c *navController) GetLatestNav(ctx *gin.Context) {

	var (
		nav model.NavDate
		req dto.FundIdenDTO
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := c.service.QueryLatestNavByFundID(&nav, req.FundID)

	if err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		response := &nav
		ctx.JSON(http.StatusOK, response)
	}
}

type pastNavByID struct {
	FundID string `form:"fid"`
	Range  string `form:"range"`
}

type pastNavByFundCode struct {
	FundCode string `form:"code"`
	Range    string `form:"range"`
}
