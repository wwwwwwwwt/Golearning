/*
 * @Author: zzzzztw
 * @Date: 2023-05-31 10:57:12
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-31 11:08:07
 * @FilePath: /Golearning/example/channel/limits.go
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}

	close(requests)

	// 200ms处理一次请求
	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println(req)
	}

	//有一定处理爆发请求的能力，然后每200ms处理一次

	burstylimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstylimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstylimiter <- t
		}
	}()

	burstyrequest := make(chan int, 5)

	for i := 0; i < 5; i++ {
		burstyrequest <- i
	}

	close(burstyrequest)

	for req := range burstyrequest {
		<-burstylimiter
		fmt.Println(req, time.Now())
	}

}
