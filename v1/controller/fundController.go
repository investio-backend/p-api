package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
	"gopkg.in/olahol/melody.v1"
)

type FundController interface {
	GetFundByID(ctx *gin.Context)
	HandleSocket(ctx *gin.Context)
}

type controller struct {
	service service.FundService
	melody  *melody.Melody
}

func NewFundController(service service.FundService, m *melody.Melody) FundController {
	// m.HandleConnect(func(s *melody.Session) {
	// 	fmt.Println("Connected: " + s.Request.Host)
	// }
	m.HandleConnect(service.HandleConnect)

	return &controller{
		service: service,
		melody:  m,
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

func (c *controller) HandleSocket(ctx *gin.Context) {
	c.melody.HandleRequest(ctx.Writer, ctx.Request)
}
