package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/request"
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

// NewFundController - A constructor of FundController
func NewFundController(service service.FundService, m *melody.Melody) FundController {
	c := &controller{
		service: service,
		melody:  m,
	}

	c.initWebsocket()

	return c
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

func (c *controller) initWebsocket() {
	c.melody.HandleConnect(c.service.HandleWsConnect)
	c.melody.HandleMessage(c.handleWsMessage)
}

func (c *controller) handleWsMessage(s *melody.Session, query []byte) {
	var (
		reqJSON  request.FundJSON
		response []byte
	)
	json.Unmarshal(query, &reqJSON)
	reqTopic := strings.ToLower(reqJSON.Topic)

	if reqTopic == "search" {
		funds, err := c.service.SearchFund(reqJSON.Data)
		if err != nil {
			panic(err)
		}
		response, _ = json.Marshal(funds)
		fmt.Println(reqJSON.Topic, reqJSON.Data)
	}

	c.melody.BroadcastFilter(response, func(q *melody.Session) bool {
		return q.Request.URL.Path == s.Request.URL.Path
	})
}
