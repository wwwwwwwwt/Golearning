/*
 * @Author: zzzzztw
 * @Date: 2023-04-21 23:41:57
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-22 00:03:46
 * @FilePath: /Golearning/Gostudy/class11_slicecap.go
 */
package main

import (
	"fmt"
)

func main() {

	a := make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(a), cap(a), a)

	a = append(a, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(a), cap(a), a)

	a = append(a, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(a), cap(a), a)

	a = append(a, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(a), cap(a), a)

	//s2 := make([]int, 3)

	//copy(s2, a[3:6])

	s2 := a[3:6]
	s2[0] = 100

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s2), cap(s2), s2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(a), cap(a), a)
}
