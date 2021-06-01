package amqp

import (
	"fmt"
	"log"
	"serqol/go-demo/logging"
	"serqol/go-demo/utils"
	"sync"

	"github.com/streadway/amqp"
)

type Amqp struct {
	connection     *amqp.Connection
	channel        *amqp.Channel
	configuration  map[string]string
	confirmChannel chan amqp.Confirmation
}

func (instance *Amqp) Publish(body []byte) {
	exchange, queue := instance.configuration["exchange"], instance.configuration["queue"]
	if err := instance.channel.Publish(
		exchange,
		queue,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "application/json",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Persistent,
			Priority:        0,
		},
	); err != nil {
		log.Fatal(err)
	}
	// implement delivery tag tracking, for now it doesn't really know which message it acks
	go confirmOne(instance.confirmChannel)
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
		exchange, exchangeType, queue := configuration["exchange"], configuration["exchange_type"], configuration["queue"]
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
		queueDeclare(channel, queue)
		bind(channel, queue, exchange)
		if err := channel.Confirm(false); err != nil {
			log.Fatal(err)
		}
		confirmationChannel := make(chan amqp.Confirmation, 10000)
		confirms := channel.NotifyPublish(confirmationChannel)
		instances[configurationHash] = &Amqp{connection, channel, configuration, confirms}
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

func queueDeclare(channel *amqp.Channel, queue string) {
	_, err := channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func bind(channel *amqp.Channel, queue string, exchange string) {
	channel.QueueBind(queue, queue, exchange, false, nil)
}

func confirmOne(confirms <-chan amqp.Confirmation) {
	confirmed := <-confirms
	if confirmed.Ack {
		logging.Log(fmt.Sprintf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag), nil)
		return
	}
	logging.Log(fmt.Sprintf("failed delivery of delivery tag: %d", confirmed.DeliveryTag), nil)
}
