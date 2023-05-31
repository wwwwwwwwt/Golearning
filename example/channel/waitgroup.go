/*
 * @Author: zzzzztw
 * @Date: 2023-05-31 10:43:19
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-31 10:50:56
 * @FilePath: /Golearning/example/channel/waitgroup.go
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int) {
	fmt.Println("worker ", id, "starting")
	time.Sleep(1 * time.Second)
	fmt.Println("worker ", id, "starting")
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			work(i)
		}()
	}

	wg.Wait()

}
