package example

import "fmt"

func Channels() {
	// 创建channel
	messages := make(chan string)

	// 将消息传入channel
	go func() { messages <- "ping" }()

	// 接收消息
	// 默认发送和接受操作时阻塞的
	msg := <-messages
	fmt.Println(msg)
}
