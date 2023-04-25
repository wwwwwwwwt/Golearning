/*
 * @Author: zzzzztw
 * @Date: 2023-04-25 10:47:59
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-25 12:40:55
 * @FilePath: /Golearning/Gostudy/class28_channel_select.go
 */
package main

import "fmt"

func main() {

	c := make(chan int)
	quit := make(chan int)
	go func() {
		x, y := 1, 1

		for {
			select {
			case c <- x:
				//x, y = y, x+y // 斐波那契数列
				z := x
				x = y
				y = z + y
			case <-quit:
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		data := <-c

		fmt.Println(data)
	}
	quit <- 0
}
