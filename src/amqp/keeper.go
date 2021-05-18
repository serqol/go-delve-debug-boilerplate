package amqp

import (
	"fmt"
	"log"
	"serqol/go-demo/utils"
	"sync"

	"github.com/streadway/amqp"
)

type Amqp struct {
	connection *amqp.Connection
}

var instances map[string]*Amqp
var once sync.Once

func Instance(configuration map[string]string) *Amqp {
	configurationHash := utils.GetMapHash(configuration)
	if _, ok := instances[configurationHash]; ok {
		return instances[configurationHash]
	}
	once.Do(func() {
		host, user, password, port := configuration['host'], configuration['user'], configuration['password'], configuration['port']
		connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", host, user, password, port)
		connection, err := amqp.Dial(connStr)
		if err == nil {
			err = connection.Ping()
		}
		if err != nil {
			log.Fatal(err)
		}
		instances[configurationHash] = &Amqp{connection}
	})
	return instance
}
