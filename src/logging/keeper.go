package logging

import (
	"fmt"
	"serqol/go-demo/utils"
)

func Log(message string, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["message"] = message
	body := utils.ToJson(data) + "\n"
	fmt.Printf(body)
}
