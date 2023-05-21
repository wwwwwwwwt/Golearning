/*
 * @Author: zzzzztw
 * @Date: 2023-05-22 00:03:22
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-22 00:07:48
 * @FilePath: /Golearning/example/const.go
 */
package main

import (
	"fmt"
	"math"
)

// 常数声明一个常量，常数表达式可以支持任意精度的运算
// 数值型常量没有确定的类型，知道给定某个确定的类型
// 比如显示类型转换
// 一个数字可以根据上下文的需要，比如变量赋值，函数调用自动确定类型

const s = "string"

func main() {

	fmt.Println(s)

	//常数表达式
	const d = 3e20 / 5000000

	fmt.Println(int64(d))

	//数值型常数
	const n = 50000
	fmt.Println(math.Sin(n))

}
