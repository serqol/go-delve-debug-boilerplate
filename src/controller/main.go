package controller

import (
	"serqol/go-demo/logging"

	"github.com/gin-gonic/gin"
)

type Main struct {
	Base *BaseController
}

func Instance() *Main {
	return &Main{&BaseController{}}
}

func (controller Main) Show(c *gin.Context) {
	logging.Log("tits are kewl", map[string]interface{}{
		"log": "tits",
	})
	controller.Base.render(c, gin.H{
		"title": "Hello, me",
	}, "index.html")
}
