package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/dto"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
)

// NavController manages NAV
type NavController interface {
	GetPastNavSeries(ctx *gin.Context)
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

func (c *navController) GetPastNavSeries(ctx *gin.Context) {
	var pastNav []model.NavDate
	var req dto.PastNavDTO

	fundID := ctx.Params.ByName("id")
	// TODO: Check if fundID is number

	if ctx.ShouldBind(&req) != nil {
		req.Range = "1mo"
	}

	fmt.Println(req.Range)

	err := c.service.GetPastNavByFundID(&pastNav, fundID, req.Range)
	if err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		fID, _ := strconv.ParseInt(fundID, 10, 32)
		response := model.NavSeries{
			FundID: int32(fID), Navs: pastNav,
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *navController) GetLatestNav(ctx *gin.Context) {

	var nav model.NavDate

	fundID := ctx.Params.ByName("id")
	// TODO: Validate if fundID is number

	err := c.service.QueryLatestNavByFundID(&nav, fundID)

	if err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		response := &nav
		ctx.JSON(http.StatusOK, response)
	}
}
