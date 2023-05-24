/*
 * @Author: zzzzztw
 * @Date: 2023-05-24 09:58:40
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-24 10:19:51
 * @FilePath: /Golearning/example/basego/range.go
 */
package main

import "fmt"

func main() {

	//1. 使用range 来遍历数组，获得值,
	nums := make([]int, 5)
	for i := 0; i < len(nums); i++ {
		nums[i] = i * 5
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum = ", sum)

	//2. 使用range 获取key
	for i, num := range nums {
		if num == 20 {
			fmt.Println("idx = ", i)
		}
	}

	//3. 使用range 遍历map，获得值

	mp := map[string]int{"111": 1, "222": 2}

	for key, val := range mp {
		fmt.Println("key, val = ", key, val)
	}

	//4. 也可以只遍历key

	for key := range mp {
		fmt.Println("key = ", key)
	}

	//5. 遍历字符串

	for i, c := range "go" {
		fmt.Println("i, c = ", i, string(c)) //不用string强制转换的话就是unicode码
	}

}
