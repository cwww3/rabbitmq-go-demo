- 以上示例每个任务只传递给一个consumer。 
- 在这一部分中，将向多个消费者传递一个消息。这就是所谓的“订阅/发布模式”。

- 在前面部分中，我们向队列发送消息和从队列接收消息。现在是时候在Rabbit中引入完整的消息传递模型了。
- RabbitMQ消息传递模型中的核心思想是生产者从不将任何消息直接发送到队列。实际上，生产者经常甚至根本不知道是否将消息传递到任何队列。

- 相反，生产者只能将消息发送到交换器。交换器是非常简单的东西。一方面，它接收来自生产者的消息，另一方面，将它们推入队列。交换器必须确切知道如何处理接收到的消息。它应该被附加到特定的队列吗？还是应该将其附加到许多队列中？或者它应该被丢弃。这些规则由交换器的类型定义。
- fanout类型交换器


```shell
# 查看交换器
sudo rabbitmqctl list_exchanges

# 查看绑定
sudo rabbitmqctl list_bindings
```

## 运行两个consumer
go run 6.publish_subscribe/consumer.go
go run 6.publish_subscribe/consumer.go
## 运行send.sh
bash 6.publish_subscribe/send.sh