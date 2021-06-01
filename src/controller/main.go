package controller

import (
	"io/ioutil"
	"net/http"
	"serqol/go-demo/graylog"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	go graylog.LogRaw(x)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"error":   0,
	})
}
