package amqp

import (
	"encoding/json"
	"fmt"
	"json"
	"log"
	"serqol/go-demo/utils"
	"sync"

	"github.com/streadway/amqp"
)

type Amqp struct {
	connection    *amqp.Connection
	channel       *amqp.Channel
	configuration map[string]string
}

func (instance *Amqp) Publish(data map[string]interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	exchange, queue := instance.configuration["exchange"], instance.configuration["queue"]
	if err := instance.channel.Publish(
		exchange,
		queue,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		log.Fatal(err)
	}
}

var instances map[string]*Amqp

func init() {
	instances = make(map[string]*Amqp)
}

func Publisher(configuration map[string]string) *Amqp {
	configurationHash := utils.GetMapHash(configuration)
	if _, ok := instances[configurationHash]; ok {
		return instances[configurationHash]
	}
	var once sync.Once
	once.Do(func() {
		host, user, password, port := configuration["host"], configuration["user"], configuration["password"], configuration["port"]
		exchange, exchangeType := configuration["exchange"], configuration["exchange_type"]
		connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
		connection, err := amqp.Dial(connStr)
		if err != nil {
			log.Fatal(err)
		}
		channel, err := connection.Channel()
		if err != nil {
			log.Fatal(err)
		}
		exchangeDeclare(channel, exchange, exchangeType)
		instances[configurationHash] = &Amqp{connection, channel, configuration}
	})
	return instances[configurationHash]
}

func exchangeDeclare(channel *amqp.Channel, exchange string, exchangeType string) {
	if err := channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		log.Fatal(err)
	}
}
