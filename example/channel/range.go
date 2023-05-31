/*
 * @Author: zzzzztw
 * @Date: 2023-05-28 17:11:54
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-31 09:53:38
 * @FilePath: /Golearning/example/channel/range.go
 */
package main

import "fmt"

func main() {

	buf := make(chan string, 2)

	buf <- "one"
	buf <- "two"
	close(buf) // 注意使用range遍历，不关闭管道会死锁
	// 非空管道也可以关闭
	for i := range buf {
		fmt.Println(i)
	}

}
