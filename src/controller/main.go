package controller

import (
	"github.com/gin-gonic/gin"
)

type Main struct {
	Base *BaseController
}

func Instance() *Main {
	return &Main{&BaseController{}}
}

func (controller Main) Show(c *gin.Context) {
	controller.Base.render(c, gin.H{
		"title": "Hello, me",
	}, "index.html")
}
