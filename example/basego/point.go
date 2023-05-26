/*
 * @Author: zzzzztw
 * @Date: 2023-05-25 09:57:12
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-25 10:02:38
 * @FilePath: /Golearning/example/basego/point.go
 */
package main

import "fmt"

func zeroval(num int) {
	fmt.Println("&valnum = ", &num)
	num = 0
}

func zerop(num *int) {
	fmt.Println("&pnum = ", num)
	*num = 0
}

func main() {
	i := 1
	fmt.Println("i = ", i, "&i = ", &i)

	zeroval(i)
	fmt.Println("i = ", i, "&i = ", &i)

	zerop(&i)
	fmt.Println("i = ", i, "&i = ", &i)
}
