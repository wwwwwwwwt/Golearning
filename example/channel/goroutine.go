/*
 * @Author: zzzzztw
 * @Date: 2023-05-26 22:02:34
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-26 22:18:07
 * @FilePath: /Golearning/example/channel/goroutine.go
 */
/*
 * @Author: zzzzztw
 * @Date: 2023-05-26 22:02:34
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-26 22:17:53
 * @FilePath: /Golearning/example/channel/goroutine.go
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var w sync.WaitGroup

	w.Add(1)
	go func(str string) {
		defer w.Done()
		fmt.Println(str)
		time.Sleep(3 * time.Second)
		//	fmt.Println("after 3 seconds")

	}("something")

	w.Add(1)

	go func(num int) {
		defer w.Done()
		fmt.Println(num)
	}(10)

	w.Wait()

	fmt.Println("all done")

}
