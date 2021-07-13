- 在一个有两个consumer的情况下，当所有的奇数消息都是重消息而偶数消息都是轻消息时，一个consumer将持续忙碌，而另一个consumer几乎不做任何工作
- 这是因为RabbitMQ只是在消息进入队列时发送消息。它不考虑消费者未确认消息的数量。只是盲目地向消费者发送信息。

- 为了避免这种情况，我们可以将预取计数设置为1。这告诉RabbitMQ不要一次向一个consumer发出多个消息。 
- 或者，换句话说，在处理并确认前一条消息之前，不要向consumer发送新消息。相反，它将把它发送给下一个不忙的consumer。
