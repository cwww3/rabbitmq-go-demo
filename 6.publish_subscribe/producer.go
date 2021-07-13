package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	common "rabbitmq-demo"
)

func main() {
	// 1. 尝试连接RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	common.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	// 3. 声明交换器
	// 声明队列是幂等的——仅当队列不存在时才创建。
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	body := common.BodyFrom(os.Args)
	// 4.将消息发布到声明的交换器
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent, // 消息持久
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	common.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
