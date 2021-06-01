package main

import (
	"serqol/go-demo/controller"
	"serqol/go-demo/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(splunkLogFormatter), gin.Recovery())
	gin.SetMode(gin.DebugMode)
	// basePath, _ := os.Getwd()
	// router.LoadHTMLGlob(basePath + "/templates/*")
	router.POST("/", controller.Show)
	router.Run(utils.GetEnv("SOCKET", "localhost:8888"))
}

var splunkLogFormatter = func(param gin.LogFormatterParams) string {
	if param.Latency > time.Minute {
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return utils.ToJson(map[string]interface{}{
		"time":     param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		"latency":  float64(param.Latency) / float64(1000000),
		"clientIp": param.ClientIP,
		"path":     param.Path,
		"error":    param.ErrorMessage,
	}) + "\n"
}
