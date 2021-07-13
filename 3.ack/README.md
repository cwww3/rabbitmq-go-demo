## 启动两个消费者
- go run workqueue/consumer.go
- go run workqueue/consumer.go
## 执行send.sh多次发送消息
bash workqueue/send.sh

## 终止其中一个consumer.go

### 结果
所有未确认的消息都将被重新发送给另一个consumer


### ACK
忘记确认是一个常见的错误。这是一个简单的错误，但后果是严重的。当你的客户机退出时，消息将被重新传递（这看起来像随机重新传递），但是RabbitMQ将消耗越来越多的内存，因为它无法释放任何未确认的消息。

```shell
# 检查是否忘记ACK
sudo rabbitmqctl list_queues name messages_ready messages_unacknowledged
```