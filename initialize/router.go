package initialize

import (
	"github.com/gin-gonic/gin"
)

func NewRouters() *gin.Engine {
	var router = gin.Default()
	//xRouter := router.Group("a")
	//utils.Unuse(v1.ApiGroup)
	//accountApi := v1.ApiGroupApp.AuthApiGroup.AccountApi
	//{
	//
	//}
	//xRouter.GET("a", .sayHello)
	return router
}
