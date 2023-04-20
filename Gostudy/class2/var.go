/*
 * @Author: zzzzztw
 * @Date: 2023-04-21 00:06:06
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 00:21:12
 * @FilePath: /Golearning/Gostudy/class2/var.go
 */
package main

import (
	"fmt"
)

// 声明全局变量 方法123是一样的
var gobal int = 1100

// ：= 只能用在函数体内，不能声明全局变量

func main() {

	// 方法1， 不给初始化变量默认为0
	var a int

	fmt.Println("a = ", a)

	// 方法2， 给初始化值
	var b int = 100

	fmt.Println("b = ", b)

	//方法3， 靠值自动推断
	var c = 110

	fmt.Printf("type of c is %T, val is %d\n", c, c)

	//方法4， 省去var关键字, 比较常用，但不能声明全局变量

	d := "abcd"

	fmt.Printf("type of c is %T, val is %s\n", d, d)

	// 多变量初始化
	xx, yy := 100, "abc"
	var zz, ee = 100, "abc"
	fmt.Printf("type of xx is %T, val is %d\n", xx, xx)
	fmt.Printf("type of yy is %T, val is %s\n", yy, yy)
	fmt.Printf("type of zz is %T, val is %d\n", zz, zz)
	fmt.Printf("type of ee is %T, val is %s\n", ee, ee)
}
