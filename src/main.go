package main

import (
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
	mainController = controller.Instance()
	router.LoadHTMLGlob(basePath + "/templates/*")
	router.GET("/", mainController.Show)
	router.Run()
}
