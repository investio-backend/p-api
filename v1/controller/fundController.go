package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
)

type FundController interface {
	GetFundByID(ctx *gin.Context)
}

type controller struct {
	service service.FundService
}

// NewFundController - A constructor of FundController
func NewFundController(service service.FundService) FundController {
	return &controller{
		service: service,
	}
}

func (c *controller) GetFundByID(ctx *gin.Context) {
	code := ctx.Params.ByName("id")
	var fund model.Fund

	err := c.service.GetFundByID(&fund, code)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, fund)
	}
}
