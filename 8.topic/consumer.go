package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
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

	// 声明交换器
	err = ch.ExchangeDeclare(
		"logs_topic", // name 交换器名字
		"topic",      // type 发布订阅
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	// 声明队列
	q, err := ch.QueueDeclare(
		"",    // name 传递一个空字符串作为队列名称时，将使用随机生成的名称创建一个非持久队列：
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}
	// 建立多个绑定关系
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_topic", s)
		err = ch.QueueBind(
			q.Name,       // queue name
			s,            // routing key
			"logs_topic", // exchange
			false,
			nil)
		common.FailOnError(err, "Failed to bind a queue")
	}

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
