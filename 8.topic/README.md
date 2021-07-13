- topic类型交换器
- 发送到topic交换器的消息不能具有随意的routing_key——它必须是单词列表，以点分隔。
- 绑定键也必须采用相同的形式。但是，绑定键有两个重要的特殊情况： *可以代替一个单词。 ＃可以替代零个或多个单词。
## 接收
- 接收所有消息
  go run 8.topic/consumer.go "#"
  
- 要从kern接收所有日志
  go run 8.topic/consumer.go "kern.*"
    
- 只接收critical日志
  go run 8.topic/consumer.go "*.critical"
  
- 绑定多个
  go run 8.topic/consumer.go "kern.*" "*.critical"
  
## 发送
go run 8.topic/producer.go "kern.critical" "A critical kernel error"