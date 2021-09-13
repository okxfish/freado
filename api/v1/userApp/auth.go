package userApp

import (
	"freado/model/response"
	"github.com/gin-gonic/gin"
)

type AuthApi struct {
}

func (receiver *AuthApi) SignIn(c *gin.Context) {
	response.OkWithMessage("asd", c)
}

func (receiver *AuthApi) SignUp(c *gin.Context) {
	response.OkWithMessage("asd", c)
}
