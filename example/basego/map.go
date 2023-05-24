/*
 * @Author: zzzzztw
 * @Date: 2023-05-24 09:30:25
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-24 09:58:30
 * @FilePath: /Golearning/example/map.go
 */
package main

import "fmt"

func main() {
	//1. 创建一个新map
	mp := make(map[string]int)

	//2. 使用map
	mp["ztw"] = 24 * 15
	mp["wgl"] = 20 * 15

	//查看下结果

	fmt.Println("mp = ", mp)

	//3. 遍历map
	for key, val := range mp {
		fmt.Println("key = ", key, "val = ", val)
	}
	// 不关心键的话可以用_占位
	for _, val := range mp {
		fmt.Println("val = ", val)
	}
	// 也可以只遍历键
	for key := range mp {
		fmt.Println("key = ", key)
	}

	//4. 查看map有没有这个键
	if val, ok := mp["ztw"]; ok {
		fmt.Println(val)
	}

	//5. 删除键
	delete(mp, "wgl")
	val, ok := mp["wgl"]
	fmt.Println("val = ", val, "ok = ", ok) // 0 false

}
