/*
 * @Author: zzzzztw
 * @Date: 2023-05-23 08:41:08
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-23 08:46:15
 * @FilePath: /Golearning/example/array.go
 */
package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("a: ", a)

	a[4] = 100

	fmt.Println("a[4]: ", a[4])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	var twoD [2][3]int

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = i * j
		}
	}

	fmt.Println("twoD: ", twoD)
}
