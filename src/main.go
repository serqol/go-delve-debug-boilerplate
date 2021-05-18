package main

import (
	"os"
	"serqol/go-demo/controller"
	"serqol/go-demo/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var mainController *controller.Main

func main() {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(defaultLogFormatter), gin.Recovery())
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

var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	if param.Latency > time.Minute {
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return utils.ToJson(map[string]interface{}{
		"time":     param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		"latency":  param.Latency,
		"clientIp": param.ClientIP,
		"path":     param.Path,
		"error":    param.ErrorMessage,
	})
}
