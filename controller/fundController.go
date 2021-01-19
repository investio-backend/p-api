package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/investio/backend/api/model"
	"gitlab.com/investio/backend/api/service"
)

func GetFundByCode(c *gin.Context) {
	code := c.Params.ByName("code")
	var fund model.Fund

	err := service.GetFundByFundCode(&fund, code)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, fund)
	}
}
