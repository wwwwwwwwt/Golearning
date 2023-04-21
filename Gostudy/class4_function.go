/*
 * @Author: zzzzztw
 * @Date: 2023-04-21 10:19:28
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 13:24:26
 * @FilePath: /Golearning/Gostudy/class4_function.go
 */
package main

import (
	"fmt"
)

// func 函数名（参数 参数类型） 返回值类型 匿名
func foo1(a string, b int) (int, int) {

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	d := 200
	return c, d

}

// func 函数名（参数 参数类型） 返回值类型 有名
func foo2(a string, b int) (c int, d int) {

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// c d 的作用域空间时foo2 的整个空间
	fmt.Println("c = ", c) // 默认为0
	fmt.Println("d = ", d) // 默认为0
	c = 100
	d = 200
	return c, d

}

func swap(a string, b string) (string, string) {
	return b, a
}

func main() {

	a, b := "abc", 20

	d, e := foo2(a, b)

	fmt.Println("d = ", d)
	fmt.Println("e = ", e)

	x, y := "abc", "bcd"

	x, y = swap(x, y)

	fmt.Println("x = ", x, "y = ", y)

}
