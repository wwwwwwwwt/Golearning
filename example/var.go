/*
 * @Author: zzzzztw
 * @Date: 2023-05-21 23:57:53
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-21 23:59:50
 * @FilePath: /Golearning/example/var.go
 */
package main

import "fmt"

func main() {
	// 1. var声明一个或多个变量

	var a = "inital"
	fmt.Println(a)
	var b, c int = 10, 20
	fmt.Println("b, c = ", b, c)

	// 2. 没有初始值会是0

	var e int
	fmt.Println(e)

	// 3. := 直接赋值

	f := "golang"

	fmt.Println(f)

}
