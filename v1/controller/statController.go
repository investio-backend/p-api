package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/dto"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
)

type StatController interface {
	GetTopReturn(ctx *gin.Context)
	GetStatInfo(ctx *gin.Context)
}

type statController struct {
	statService service.StatService
}

// NewStatController - A constructor of FundController
func NewStatController(service service.StatService) StatController {
	return &statController{
		statService: service,
	}
}

func (c *statController) GetTopReturn(ctx *gin.Context) {
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
		if err := c.statService.FindTopStat1Y(&statRes, queryStr.Cat, queryStr.Amc); err != nil {
			// fmt.Println("Return Err: ", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			// fmt.Println("Res: ", statRes)
			ctx.JSON(http.StatusOK, statRes)
		}
	} else if strings.ToLower(queryStr.Range) == "6m" {
		var statRes []model.Stat_6M
		if err := c.statService.FindTopStat6M(&statRes, queryStr.Cat, queryStr.Amc); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			ctx.JSON(http.StatusOK, statRes)
		}
	} else {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
	}
}

func (c *statController) GetStatInfo(ctx *gin.Context) {
	fundID := ctx.Params.ByName("fundID")
	var statResult model.StatFundResponse

	if err := c.statService.GetStatByFundID(&statResult, fundID); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, statResult)
}
