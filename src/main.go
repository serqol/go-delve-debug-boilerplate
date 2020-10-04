package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"serqol/go-demo/service"
)

var utils *service.Utils
var router *gin.Engine

func main() {
	services := utils.GetTestObjects(10)
	for key, object := range(services) {
		fmt.Println(string(key), object)
	}
	router.Run()
}

type Service struct {
	id int
	name string
}

func buildIndex(collection []interface{}, key string) []interface{} {
	return collection
}