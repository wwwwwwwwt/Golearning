/*
 * @Author: zzzzztw
 * @Date: 2023-05-31 09:21:59
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-31 09:34:22
 * @FilePath: /Golearning/print.go
 */
package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	A := make(chan bool, 1)
	B := make(chan bool)
	wg.Add(1)
	wg.Add(1)
	A <- true
	go func() {
		for i := 0; i < 5; i++ {
			<-A
			fmt.Println("A")
			B <- true
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			<-B
			fmt.Println("B")
			A <- true
		}
		wg.Done()
	}()

	wg.Wait()

}
