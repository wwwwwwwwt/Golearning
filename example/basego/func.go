/*
 * @Author: zzzzztw
 * @Date: 2023-05-24 10:20:49
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-25 01:01:16
 * @FilePath: /Golearning/example/basego/func.go
 */
package main

import "fmt"

func f(n int) int {
	if n < 2 {
		return n
	}

	return f(n-1) + f(n-2)
}

func main() {

	//递归调用闭包
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	a := f(10)
	fmt.Println(fib(10))
	fmt.Println(a)
}
