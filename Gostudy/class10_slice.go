/*
 * @Author: zzzzztw
 * @Date: 2023-04-21 23:29:15
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 23:39:58
 * @FilePath: /Golearning/Gostudy/class10_slice.go
 */
package main

import (
	"fmt"
)

func main() {

	//方法1 声明长度为3的切片数组

	slice1 := []int{1, 2, 3}

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 方法2 声明长度为0的切片数组,没有分配容量，可以通过make([]类型，长度)进行分配空间，值为默认值

	var slice2 []int

	slice2 = make([]int, 4)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	//方法3， 通过var 声明

	var slice3 []int = make([]int, 5)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	// 方法4， 动态直接初始化，分配空间

	slice4 := make([]int, 6)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	// 判断是否为空，就是判断是否为nil
	var a []int

	if a == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
