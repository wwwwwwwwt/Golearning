package main

import "fmt"

func main() {

	// 1. 无缓冲通道
	//只有对应的接受端准备好接受走channel内数据，才允许发送
	message := make(chan string)

	go func() {
		message <- "some things"
	}()

	msg := <-message // 阻塞等待，使两个协程同步

	fmt.Println(msg)

	// 2. 带缓冲的通道

	ch := make(chan string, 2)

	ch <- "123"
	ch <- "456"

	fmt.Println(<-ch, <-ch)
}
