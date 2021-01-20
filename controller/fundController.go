package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/model"
	"gitlab.com/investio/backend/api/service"
)

type FundController interface {
	GetFundByCode(ctx *gin.Context)
}

type controller struct {
	service service.FundService
}

func New(service service.FundService) FundController {
	return &controller{
		service: service,
	}
}

func (c *controller) GetFundByCode(ctx *gin.Context) {
	code := ctx.Params.ByName("code")
	var fund model.Fund

	err := c.service.GetFundByFundCode(&fund, code)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, fund)
	}
}
