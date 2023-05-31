/*
 * @Author: zzzzztw
 * @Date: 2023-05-31 09:54:53
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-31 09:58:21
 * @FilePath: /Golearning/example/channel/Ticker.go
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go func() {

		for {
			select {
			case <-done:
				return
			case t := <-time.After(2 * time.Second):
				fmt.Println("afetr 2s ", t)
			}
		}

	}()

	time.Sleep(10 * time.Second)
	done <- true

}
