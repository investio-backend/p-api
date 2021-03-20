package controller

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/dto"
	"gopkg.in/olahol/melody.v1"
)

// SocketController manages websocket
type SocketController interface {
	HandleSocket(ctx *gin.Context)
}

type socketController struct {
	fundController FundController
	melody         *melody.Melody
}

// NewSocketController - A contructor of SocketController
func NewSocketController(m *melody.Melody, fundController FundController) SocketController {
	c := &socketController{
		fundController: fundController,
		melody:         m,
	}

	c.initWebsocket()

	return c
}

func (c *socketController) HandleSocket(ctx *gin.Context) {
	c.melody.HandleRequest(ctx.Writer, ctx.Request)
}

func (c *socketController) initWebsocket() {
	c.melody.HandleConnect(c.handleWsConnected)
	c.melody.HandleMessage(c.handleWsMessage)
}

func (c *socketController) handleWsMessage(s *melody.Session, query []byte) {
	var (
		reqJSON  dto.SocketDTO
		response []byte
	)
	json.Unmarshal(query, &reqJSON)
	reqType := strings.ToUpper(reqJSON.Type)
	reqTopic := strings.ToLower(reqJSON.Topic)

	if reqType == "FUND" {
		if reqTopic == "search" {
			response = c.fundController.SocketSearchFund(reqJSON)
		}
	}

	if response != nil {
		c.melody.BroadcastFilter(response, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	}
}

func (c *socketController) handleWsConnected(s *melody.Session) {
	fmt.Println("Connected: " + s.Request.Host)
}
