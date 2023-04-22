/*
 * @Author: zzzzztw
 * @Date: 2023-04-22 16:54:39
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-22 17:05:30
 * @FilePath: /Golearning/Gostudy/class12_map.go
 */
package main

import "fmt"

func main() {

	// 定义mp 为map[key] val类型的map,但这样声明没有底层空间，是nil
	var mp map[string]string

	if mp == nil {
		fmt.Println("map is nil")
	}

	// 方法1：指定底层cap大小
	mp = make(map[string]string, 2) //给mp分配2个空间， 一旦满了，就会自动动态扩容

	mp["abc"] = "123"
	mp["c"] = "123"
	mp["d"] = "123"
	mp["e"] = "123"
	fmt.Printf("size = %d, map = %v\n", len(mp), mp)

	// 方法2：不指定大小，会自动扩容
	mp2 := make(map[int]string)
	mp2[1] = "abc"
	mp2[3] = "bcd"
	fmt.Printf("map = %v\n", mp2)

	// 方法3， 直接以值初始化

	mp3 := map[int]string{
		1: "abc",
		2: "cdf",
		3: "rdf",
	}

	fmt.Printf("map = %v\n", mp3)

}
