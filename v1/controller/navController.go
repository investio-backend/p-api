package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/dto"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
)

// NavController manages NAV
type NavController interface {
	GetPastNavWithAsset(ctx *gin.Context)
	GetPastNav(ctx *gin.Context)
	GetNavByDate(ctx *gin.Context)
	GetPastSetIndex(ctx *gin.Context)
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

type pastNavByFundCode struct {
	FundCode string `form:"f"`
	Range    string `form:"range"`
}

func (c *navController) GetPastNavWithAsset(ctx *gin.Context) {
	var (
		pastNav       []model.NavDate
		reqByFundCode pastNavByFundCode
		err           error
	)

	if err = ctx.ShouldBind(&reqByFundCode); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// fmt.Println(reqByID.FundID)
	if reqByFundCode.FundCode != "" {
		err = c.service.FindPastNavWithAsset(&pastNav, reqByFundCode.FundCode, reqByFundCode.Range)
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err != nil {
		// fmt.Println(err.Error())
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	if pastNav == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	response := model.NavSeries{
		FundCode: reqByFundCode.FundCode, Navs: pastNav,
	}
	ctx.JSON(http.StatusOK, response)

}

type pastNavByID struct {
	FundID string `form:"f"`
	Range  string `form:"range"`
}

func (c *navController) GetPastNav(ctx *gin.Context) {
	var (
		navList [][2]interface{}
		reqByID pastNavByID
		err     error
	)
	if err = ctx.ShouldBind(&reqByID); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// fmt.Println(reqByID.FundID)
	if err = ctx.ShouldBind(&reqByID); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// fmt.Println(reqByID.FundID)
	if reqByID.FundID != "" {
		err = c.service.FindPastNavChart(&navList, reqByID.FundID, reqByID.Range)
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err != nil {
		// fmt.Println(err.Error())
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	if len(navList) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// response := model.NavSeries{
	// 	FundID: reqByID.FundID, Navs: pastNav,
	// }
	ctx.JSON(http.StatusOK, navList)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"nav": gin.H{
	// 		"data": nav,
	// 	},
	// 	"dates": date,
	// })
}

func (c *navController) GetNavByDate(ctx *gin.Context) {
	var (
		nav model.NavDate
		req dto.QueryNavByDate
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// fmt.Println("Date: ", req.Date)

	dataDate, err := time.Parse("2006-01-02", req.Date)

	// Cannot extract date -> Get Latest NAV
	if err != nil {
		err := c.service.FindLatestNavByFundID(&nav, req.FundID)
		if err != nil {
			// fmt.Println(err.Error())
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
	} else {
		// Get NAV by date
		err := c.service.FindNAVByDate(&nav, model.Date(dataDate), req.FundID)
		if err != nil {
			// fmt.Println(err.Error())
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

	}
	if nav.Date == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, nav)
	}
}

func (c *navController) GetPastSetIndex(ctx *gin.Context) {
	var result []model.SetDatePrice

	if err := c.service.FindPastSetIndex(&result, "6mo"); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err.Error())
		return
	}

	response := model.SetSeries{
		Name: "SET",
		Navs: result,
	}
	ctx.JSON(http.StatusOK, response)
}
