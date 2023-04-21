/*
 * @Author: zzzzztw
 * @Date: 2023-04-21 22:34:38
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 22:46:51
 * @FilePath: /Golearning/Gostudy/class7_point.go
 */
package main

import (
	"fmt"
)

func changeval(a *int) {
	*a = 20
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {

	var a int = 10
	fmt.Println("a = ", a, "&a = ", &a)

	changeval(&a)

	fmt.Println("a = ", a, "&a = ", &a)

	fmt.Println("-------swap val --------")

	x, y := 100, 200

	fmt.Println("x = ", x, "y = ", y)

	swap(&x, &y)

	fmt.Println("x = ", x, "y = ", y)
}
