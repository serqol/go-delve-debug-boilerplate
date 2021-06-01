package graylog

import (
	"serqol/go-demo/amqp"
	"serqol/go-demo/utils"

	"github.com/joho/godotenv"
)

var amqpInstance *amqp.Amqp

func init() {
	godotenv.Load(".env")
	configuration := map[string]string{
		"host":          utils.GetEnv("GRAYLOG_HOST", "rabbit"),
		"user":          utils.GetEnv("GRAYLOG_USER", "guest"),
		"password":      utils.GetEnv("GRAYLOG_PASSWORD", "guest"),
		"port":          utils.GetEnv("GRAYLOG_PORT", "5672"),
		"exchange":      utils.GetEnv("GRAYLOG_EXCHANGE", "graylog"),
		"exchange_type": utils.GetEnv("GRAYLOG_EXCHANGE_TYPE", "direct"),
		"queue":         utils.GetEnv("GRAYLOG_QUEUE", "graylog"),
	}
	amqpInstance = amqp.Publisher(configuration)
}

func LogRaw(message []byte) {
	amqpInstance.Publish(message)
}

func Log(message string, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["message"] = message
	body := utils.ToJson(data)
	amqpInstance.Publish([]byte(body))
}
