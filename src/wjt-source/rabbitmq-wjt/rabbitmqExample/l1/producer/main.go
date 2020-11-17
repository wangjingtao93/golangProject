package main

import (
	config "wjt-source/rabbitmq-wjt/rabbitmqExample/l1/conf"
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial(config.RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(config.PRODUCERCNT)//配置队列大小，比如50

	for routine := 0; routine < config.PRODUCERCNT; routine++ {
		go func(routineNum int) {
			ch, err := conn.Channel()
			failOnError(err, "Failed to open a channel")
			defer ch.Close()

			q, err := ch.QueueDeclare(
				config.QUEUENAME,//producer和consumer要一致
				false,
				false,
				false,
				false,
				nil,
			)

			failOnError(err, "Failed to declare a queue")

			for i := 0; i < 65535; i++ {
				msgBody := fmt.Sprintf("Message_%d_%d", routineNum, i) //要发送的字段

				err = ch.Publish(                                      //producer是Publish， consumer是Consume
					"",     //exchange
					q.Name, //routing key
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",                  //producer和consumer要一致
						Body:        []byte(msgBody),               //填充要发送的字段
					})

				log.Printf(" [x] Sent %s", msgBody)
				failOnError(err, "Failed to publish a message")
			}

			wg.Done()
		}(routine)
	}

	wg.Wait()

	log.Println("All messages sent!!!!")
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
