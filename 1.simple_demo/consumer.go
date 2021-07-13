package main

import (
	"github.com/streadway/amqp"
	"log"
	common "rabbitmq-demo"
)

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	common.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 获取channel
	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明队列
	// 可能在发布者之前启动使用者，所以我们希望在尝试使用队列中的消息之前确保队列存在
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	// 获取接收消息的Delivery通道
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	common.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}