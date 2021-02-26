package controller

import (
	"serqol/go-demo/service"

	"github.com/gin-gonic/gin"
)

type Main struct {
	Base *BaseController
}

func Instance() *Main {
	return &Main{&BaseController{}}
}

func (controller Main) Show(c *gin.Context) {
	dbInstance := service.DatabaseInstance()
	dbInstance.Connection.Ping()
	controller.Base.render(c, gin.H{
		"title": "Hello, me",
	}, "index.html")
}
