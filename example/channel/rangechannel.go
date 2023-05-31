/*
 * @Author: zzzzztw
 * @Date: 2023-05-31 09:47:19
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-31 09:49:47
 * @FilePath: /Golearning/example/channel/rangechannel.go
 */
package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for val := range queue {
		fmt.Println(val)

	}
}
