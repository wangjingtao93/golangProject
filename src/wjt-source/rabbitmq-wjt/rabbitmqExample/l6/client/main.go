package main

import (
	config "wjt-source/rabbitmq-wjt/rabbitmqExample/l6/conf"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/streadway/amqp"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("Arguments error")
		return
	}

	conn, err := amqp.Dial(config.RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgBody := os.Args[1]

	respQueue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,
	)
	failOnError(err, "Failed to declare a response queue")

	correlationID := randomID(32)              //校验值之类的

	err = ch.Publish(
		"",               //exchange
		config.QUEUENAME, //routing key
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: correlationID,     //填充校验值
			ReplyTo:       respQueue.Name,
			Body:          []byte(msgBody),   //填充要发送的字段
		})

	log.Printf(" [x] Sent %s", msgBody)
	failOnError(err, "Failed to publish a message")

	respMsgs, err := ch.Consume(          //获取返回值
		respQueue.Name,
		"",
		true,  // auto-ack
		true,  // exclusive
		false, // noLocal
		false, // nowait
		nil,
	)

	for item := range respMsgs {
		if item.CorrelationId == correlationID {
			fmt.Println("response:", string(item.Body))
			break
		}
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}

func randomID(length int) string {
	if length <= 0 {
		return ""
	}

	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(rand.Intn(9))
	}

	return string(bytes)
}
