/*
 * @Author: zzzzztw
 * @Date: 2023-04-21 09:41:42
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 09:57:32
 * @FilePath: /Golearning/Gostudy/class3/const.go
 */
package main

import (
	"fmt"
)

// 通过const 定义枚举类型

const (
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2 // 1, 2
	c, d                      // 2, 3
	e, f                      // 4, 5

	g, h = iota * 2, iota * 3 // 6,9
	i, k                      // , 12

)

func main() {

	const length int = 8
	fmt.Printf("type of length is %T\n", length)

	fmt.Println("SHANGHAI", SHANGHAI)
	fmt.Println("SHENZHEN", SHENZHEN)

	fmt.Printf("i,k = %d %d\n", i, k)

}
