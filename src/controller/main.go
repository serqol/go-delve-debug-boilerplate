package controller

import (
	"serqol/go-demo/graylog"

	"github.com/gin-gonic/gin"
)

type Main struct {
	Base *BaseController
}

func Instance() *Main {
	return &Main{&BaseController{}}
}

func (controller Main) Show(c *gin.Context) {
	graylog.Log("data is kewl", nil)
	controller.Base.render(c, gin.H{
		"title": "Hello, me",
	}, "index.html")
}
