package core

import (
	"freado/initialize"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func RunWindowsServer() {
	// windows 运行
	// 注入路由 以及其他windows中间件

	router := initialize.NewRouters()
	//initServer("localhost:8888", router)
	router.Run()
}
