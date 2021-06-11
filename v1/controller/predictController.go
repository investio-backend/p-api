package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/investio/backend/api/v1/dto"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
)

type PredictController interface {
	GetTopPredict(ctx *gin.Context)
	AddPredict(ctx *gin.Context)
}

type predictController struct {
	predictService service.PredictService
}

func NewPredictController(predict service.PredictService) PredictController {
	return &predictController{
		predictService: predict,
	}
}

type queryTopPredict struct {
	Risk uint
}

func (c *predictController) GetTopPredict(ctx *gin.Context) {
	var (
		queryStr queryTopPredict
		results  []model.PredictBuyResponse
	)

	if ctx.ShouldBind(&queryStr) != nil {
		queryStr = queryTopPredict{
			Risk: 9,
		}
	}
	if queryStr.Risk == 0 {
		queryStr.Risk = 9
	}
	if err := c.predictService.ReadTopResult(&results, queryStr.Risk); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err.Error())
	}
	ctx.JSON(http.StatusOK, results)
}

func (c *predictController) AddPredict(ctx *gin.Context) {
	var req dto.PredictPostReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid data provided " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"reason": "Invalid data provided " + err.Error(),
		})
		return
	}

	predict := model.PredictBuy{
		FundCode: req.FundCode,
		Fiid:     req.Fiid,
		DataDate: req.DataDate.ParseTime(),
		Prob:     req.Prob,
	}

	if err := c.predictService.WriteResult(&predict); err != nil {
		log.Error("Write result failed" + err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"reason": "Write result failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, "OK")
}
