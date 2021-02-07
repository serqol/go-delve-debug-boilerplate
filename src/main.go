package main

import (
	"net/http"
	"os"
	"serqol/go-demo/controller"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var mainController *controller.Main

func main() {
	router = gin.Default()
	gin.SetMode(gin.DebugMode)
	basePath, err := os.Getwd()
	if err != nil {
		// do nothing
	}
	router.LoadHTMLGlob(basePath + "/templates/*")
	router.GET("/", mainController.Show)
	router.Run()
}

// TODO: not mine
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
