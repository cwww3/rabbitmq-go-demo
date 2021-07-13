- 在本项目中，将添加一个特性-我能够只订阅消息的一个子集。例如，我们将只能将关键错误消息定向到日志文件（以节省磁盘空间），同时仍然能够在控制台上打印所有日志消息。

- direct类型交换器
- 消息进入其binding key与消息的routing key完全匹配的队列。

- 屏幕上查看warning和err级别的日志消息
go run 7.router/consumer.go warning error
- 屏幕上查看所有日志消息
go run 7.router/consumer.go info warning error
## 运行send.sh
bash  7.router/send.sh