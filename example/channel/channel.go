/*
 * @Author: zzzzztw
 * @Date: 2023-05-27 09:06:48
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-27 16:42:17
 * @FilePath: /Golearning/example/channel/channel.go
 */
package main

import (
	"fmt"
	"time"
)

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

	// 3. 通道同步

	done := make(chan bool)
	go func(done chan bool) {
		fmt.Println("working...")
		time.Sleep(1 * time.Second)
		fmt.Println("finish working")
		done <- true
	}(done)

	<-done

}
