package graylog

import (
	"serqol/go-demo/amqp"
	"serqol/go-demo/utils"
)

var amqpInstance *amqp.Amqp

func init() {
	configuration := map[string]string{
		"host":          utils.GetEnv("GRAYLOG_HOST", "rabbit"),
		"user":          utils.GetEnv("GRAYLOG_USER", "guest"),
		"password":      utils.GetEnv("GRAYLOG_PASSWORD", "guest"),
		"port":          utils.GetEnv("GRAYLOG_PORT", "5672"),
		"exchange":      utils.GetEnv("GRAYLOG_EXCHANGE", "graylog"),
		"exchange_type": utils.GetEnv("GRAYLOG_EXCHANGE_TYPE", "direct"),
	}
	amqpInstance = amqp.Publisher(configuration)
}

func log(info string, data map[string]interface{}) {

}
