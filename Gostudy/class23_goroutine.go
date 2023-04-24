/*
 * @Author: zzzzztw
 * @Date: 2023-04-24 19:43:19
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-24 20:58:14
 * @FilePath: /Golearning/Gostudy/class23_goroutine.go
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

func foo() {
	fmt.Println("hhhhhhh")
}

func main() {

	// 匿名函数执行
	go func() {
		defer fmt.Println("A defer")

		func() {
			defer fmt.Println("B defer")
			runtime.Goexit() //退出当前协程，直接退出，到执行defer那一步 / return是退出当前函数，还会执行上层函数体的后续
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	// 有名函数执行
	go foo()

	for {
		time.Sleep(1 * time.Second)
	}
}
