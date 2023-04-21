/*
 * @Author: zzzzztw
 * @Date: 2023-04-21 23:04:32
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 23:26:22
 * @FilePath: /Golearning/Gostudy/class9_slice.go
 */
package main

import (
	"fmt"
)

// 数组传参

func printarray(a []int) {

	for i := 0; i < len(a); i++ {
		a[i] = i + 10
	}
	for _, val := range a {
		fmt.Println(val)
	}
}

func main() {

	// 固定长度数组
	//var myarray [10]int

	myarray2 := [10]int{1, 2, 3, 4}

	/*for i := 0; i < len(myarray2); i++ {
		fmt.Println(myarray2[i])
	}*/

	printarray(myarray2[:])

	for _, val := range myarray2 {
		fmt.Print(val)
	}

	//查看数组数据类型
	fmt.Printf("myarray2 type is %T\n", myarray2)

	fmt.Println("-------------------")

	myarray3 := []int{1, 2, 3, 4}

	printarray(myarray3)

}
