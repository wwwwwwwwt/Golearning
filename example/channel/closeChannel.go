/*
 * @Author: zzzzztw
 * @Date: 2023-05-28 16:19:47
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-28 17:10:14
 * @FilePath: /Golearning/example/channel/closeChannel.go
 */
package main

import "fmt"

func main() {

	// 如果确定不再向管道内继续发送信息，就应该关闭管道

	buf := make(chan int, 10)
	done := make(chan bool)
	go func() {
		for {
			if val, ok := <-buf; ok {
				fmt.Println(val)
			} else {
				fmt.Println("chan is close")
				done <- true
				break
			}
		}
	}()

	for i := 0; i < 10; i++ {
		buf <- i
	}

	close(buf)
	<-done
}
