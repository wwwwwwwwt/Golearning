/*
 * @Author: zzzzztw
 * @Date: 2023-05-27 16:53:19
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-28 16:17:26
 * @FilePath: /Golearning/example/channel/select.go
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	// 1. select 会随机检测一个通道是否准备好，准备好的话就执行对应case语句然后退出，否则阻塞接着检测,

	// ch := make(chan string)
	// ch2 := make(chan string)

	// go func(ch chan<- string) {

	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		ch <- "123"
	// 	}
	// }(ch)

	// go func(ch2 chan<- string) {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		ch2 <- "234"
	// 	}

	// }(ch2)

	// for {
	// 	select {
	// 	case <-ch:
	// 		fmt.Println("ch already")
	// 	case <-ch2:
	// 		fmt.Println("ch2 already")
	// 	}
	// }

	// 场景2： 竞争选举, 多个通道有一个满足条件就可以读取，然后可以竞选成功
	// s := rand.NewSource(time.Now().UnixMicro()) //初始化随机种子
	// r := rand.New(s)
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// ch3 := make(chan int)
	// go func(ch chan<- int) {
	// 	times := r.Intn(5) + 1
	// 	time.Sleep(time.Duration(times) * time.Second)
	// 	ch <- times
	// }(ch1)
	// go func(ch chan<- int) {
	// 	times := r.Intn(5) + 1
	// 	time.Sleep(time.Duration(times) * time.Second)
	// 	ch <- times
	// }(ch2)
	// go func(ch chan<- int) {
	// 	times := r.Intn(5) + 1
	// 	time.Sleep(time.Duration(times) * time.Second)
	// 	ch <- times
	// }(ch3)
	// select {
	// case i := <-ch1:
	// 	fmt.Println("ch1 收到数据", i)
	// case j := <-ch2:
	// 	fmt.Println("ch2 收到数据", j)
	// case m := <-ch3:
	// 	fmt.Println("ch3 收到数据", m)
	// }

	// 场景3：定时器超时功能
	// 	str := make(chan string)

	// 	go func(ch chan<- string) {
	// 		for i := 1; i < 7; i++ {
	// 			time.Sleep(time.Duration(i) * time.Second)
	// 			ch <- "123"
	// 		}

	// 	}(str)

	// 	for {
	// 		select {
	// 		case i := <-str:
	// 			fmt.Println("receive: ", i)
	// 		case <-time.After(3 * time.Second):
	// 			fmt.Println("time out")
	// 			goto end
	// 		}
	// 	}
	// end:

	// 场景4：判断buf是否满了，满了就default等待
	// buf := make(chan int, 5)

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	for {
	// 		<-buf
	// 		time.Sleep(3 * time.Second)
	// 	}
	// }()

	// for {
	// 	select {
	// 	case buf <- 1:
	// 		fmt.Println("add successful")
	// 		time.Sleep(1 * time.Second)
	// 	default:
	// 		fmt.Println("buf fully. wait...")
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }

	// 场景5： 阻塞main函数，注意，一定要有活动的goroutine否则会deadlock

	buf := make(chan int)

	go func() {
		for {
			buf <- 1
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println(<-buf)
		}
	}()

	select {}
}
