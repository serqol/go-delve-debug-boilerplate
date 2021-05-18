package controller

import (
	"serqol/go-demo/database"

	"github.com/gin-gonic/gin"
)

type Main struct {
	Base *BaseController
}

func Instance() *Main {
	return &Main{&BaseController{}}
}

func (controller Main) Show(c *gin.Context) {
	err := database.Instance().Connection.Ping()
	if err != nil {
		panic("Database unreachable")
	}
	controller.Base.render(c, gin.H{
		"title": "Hello, me",
	}, "index.html")
}
