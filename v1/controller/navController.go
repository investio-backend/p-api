package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/request"
	"gitlab.com/investio/backend/api/v1/service"
)

type NavController interface {
	GetPastNavSeries(ctx *gin.Context)
}

type navController struct {
	service service.NavService
}

func NewNavController(service service.NavService) NavController {
	return &navController{
		service: service,
	}
}

func (c *navController) GetPastNavSeries(ctx *gin.Context) {
	var pastNav []model.NavDate
	var req request.PastNavRequest

	code := ctx.Params.ByName("code")

	if ctx.ShouldBind(&req) != nil {
		req.Range = "1mo"
	}

	fmt.Println(req.Range)

	err := c.service.GetPastNavByFundCode(&pastNav, code, req.Range)
	if err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		response := model.NavSeries{
			FundCode: code, Navs: pastNav,
		}
		ctx.JSON(http.StatusOK, response)
	}
}
