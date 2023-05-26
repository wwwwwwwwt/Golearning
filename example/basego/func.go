/*
 * @Author: zzzzztw
 * @Date: 2023-05-24 10:20:49
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-25 09:45:31
 * @FilePath: /Golearning/example/basego/func.go
 */
package main

import "fmt"

func p(a, b, c int) int { // 普通函数
	return a + b + c
}

func p2(a, b int) (int, int) {
	return a, b
}

func b() func() int { // 闭包函数
	i := 0
	return func() int {
		i++
		return i
	}
}

func f(n int) int { //普通递归
	if n < 2 {
		return n
	}

	return f(n-1) + f(n-2)
}

func sum(nums ...int) int { //变参函数
	var a = 0
	for _, val := range nums {
		a += val
	}
	return a
}

func main() {

	//普通函数和多返回值
	fmt.Println(p(1, 2, 3))
	fmt.Println(p2(1, 2))

	//匿名函数递归调用闭包，必须在之前先声明遍历再把匿名函数赋值给变量
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

	// 闭包函数, 闭包中的变量拥有记忆，生命周期和闭包相同，直到再次被创建
	increase := b()
	fmt.Println(increase()) //1
	fmt.Println(increase()) //2
	fmt.Println(increase()) //3

	increase = b()          //再次赋值
	fmt.Println(increase()) //1

	//变参函数
	fmt.Println(sum(1, 2, 3, 4, 5))

	sli := make([]int, 3)
	sli[0] = 5
	sli[1] = 6
	sli[2] = 7
	//将数组传入变参函数
	fmt.Println(sum(sli...))
}
