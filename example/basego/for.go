/*
 * @Author: zzzzztw
 * @Date: 2023-05-22 00:09:12
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-22 00:13:36
 * @FilePath: /Golearning/example/for.go
 */
package main

import "fmt"

func main() {

	// 1. 类似while
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// 2. 正常for循环
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	// 3. 支持continue 和 break

	for k := 0; k < 5; k++ {
		if k == 1 {
			fmt.Println(k)
			continue
		} else if k == 3 {
			fmt.Println("break in ", k)
		}
	}
}
