/*
 * @Author: zzzzztw
 * @Date: 2023-04-20 23:43:45
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 00:01:59
 * @FilePath: /Golearning/Gostudy/class1/hello.go
 */
package main // 程序的包名

import (
	"fmt"
	"time"
)

func main() {

	//语句加不加分号无所谓，一般推荐不加
	//go语言要求{}的左括号与函数名在同一行，如上所写
	fmt.Println("hello world")
	time.Sleep(1 * time.Second)
}
