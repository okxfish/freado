package userApp

import (
	"freado/model/response"
	"github.com/gin-gonic/gin"
)

type AccountApi struct {
}

func (receiver *AccountApi) Forgot(c *gin.Context) {
	response.OkWithMessage("asd", c)
}
