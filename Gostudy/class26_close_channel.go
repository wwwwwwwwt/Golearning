/*
 * @Author: zzzzztw
 * @Date: 2023-04-25 10:08:01
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-25 10:26:15
 * @FilePath: /Golearning/Gostudy/class26_close_channel.go
 */
package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			c <- i
		}

		close(c) // close 关闭一个管道，不写这行运行会报死锁的错误

	}()

	for {

		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main finished")

}
