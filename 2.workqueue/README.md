## 启动两个消费者
- go run workqueue/consumer.go
- go run workqueue/consumer.go
## 执行send.sh多次发送消息
bash workqueue/send.sh

### 结果
默认情况下，RabbitMQ将按顺序将每个消息发送给下一个消费者。平均而言，每个消费者都会收到相同数量的消息。这种分发消息的方式称为轮询。