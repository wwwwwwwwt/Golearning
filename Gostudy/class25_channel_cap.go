/*
 * @Author: zzzzztw
 * @Date: 2023-04-25 09:51:23
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-25 10:05:55
 * @FilePath: /Golearning/Gostudy/class25_channel_cap.go
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)

	fmt.Println("len c = ", len(c), "cap c =", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		c <- 666
		c <- 777
		c <- 888
		c <- 999
	}()

	time.Sleep(1 * time.Second)

	num1 := <-c
	num2 := <-c
	num3 := <-c

	fmt.Println("1 :", num1, "2 :", num2, "3 :", num3)
	time.Sleep(1 * time.Second)
}
