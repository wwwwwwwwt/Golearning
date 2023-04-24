/*
 * @Author: zzzzztw
 * @Date: 2023-04-24 20:58:26
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-24 21:48:43
 * @FilePath: /Golearning/Gostudy/class24_use_channel.go
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine is finish") // defer 会在c <-666之后执行
		c <- 666                                 // 在此处就把值通过管道传递给了主线程num，在此处同步，然后接下来继续执行下面的
		time.Sleep(1 * time.Second)
	}()

	if num := <-c; num != 0 {
		fmt.Println("num is ", num)
	}
	//fmt.Println("num is", num)

	time.Sleep(2 * time.Second)
}
