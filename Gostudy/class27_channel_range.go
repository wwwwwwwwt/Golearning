/*
 * @Author: zzzzztw
 * @Date: 2023-04-25 10:30:11
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-25 10:43:04
 * @FilePath: /Golearning/Gostudy/class27_channel_range.go
 */
package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}
		close(c)
	}()

	for data := range c { // 只要channel中有数据，就阻塞尝试取出，直到channel被关闭。channel一直不关闭会导致死锁
		fmt.Println(data)
	}

}
