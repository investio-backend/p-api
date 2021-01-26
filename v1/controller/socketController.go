package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/v1/dto"
	"gitlab.com/investio/backend/api/v1/service"
	"gopkg.in/olahol/melody.v1"
)

type SocketController interface {
	HandleSocket(ctx *gin.Context)
}

type socketController struct {
	fundService service.FundService
	melody      *melody.Melody
}

// NewSocketController - A contructor of SocketController
func NewSocketController(m *melody.Melody, fundService service.FundService) SocketController {
	c := &socketController{
		fundService: fundService,
		melody:      m,
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
			funds, err := c.fundService.SearchFund(reqJSON.Data)
			if err != nil {
				// panic(err)
				log.Fatal(err)
				errResponse := &dto.SocketDTO{
					Type:  "ERROR",
					Topic: "Database",
					Data:  err.Error(),
				}
				response, err = json.Marshal(errResponse)
				if err != nil {
					log.Fatal("Marshall DB Fail:" + err.Error())
				}
				// response = []byte("Failed: Database " + err.Error())
			} else {
				response, _ = json.Marshal(funds)
				fmt.Println(reqJSON.Topic, reqJSON.Data)
			}
		}
	}

	c.melody.BroadcastFilter(response, func(q *melody.Session) bool {
		return q.Request.URL.Path == s.Request.URL.Path
	})
}

func (c *socketController) handleWsConnected(s *melody.Session) {
	fmt.Println("Connected: " + s.Request.Host)
}
