/*
 * @Author: zzzzztw
 * @Date: 2023-05-27 16:47:52
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-27 16:51:52
 * @FilePath: /Golearning/example/channel/channel_direction.go
 */
package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "msg")
	pong(pings, pongs)

	fmt.Println(<-pongs)

}
