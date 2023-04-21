/*
 * @Author: zzzzztw
 * @Date: 2023-04-21 22:53:09
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 23:00:10
 * @FilePath: /Golearning/Gostudy/class8_defer.go
 */
package main

import (
	"fmt"
)

func foo(a int) int {
	fmt.Println("main end ", a)
	return 0
}

func foo2() int {
	defer foo(1)
	return foo(2)
}

func main() {
	/*defer foo(1)
	defer foo(2)
	defer foo(3) */

	foo2()

}
