package controller

import "github.com/gin-gonic/gin"

type Main struct {
	base BaseController
}

func (controller *Main) Show(c *gin.Context) {
	controller.base.render(c, gin.H{
		"title":   "Hello, me",
	}, "index.html")
}