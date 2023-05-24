/*
 * @Author: zzzzztw
 * @Date: 2023-05-23 08:46:31
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-23 10:19:56
 * @FilePath: /Golearning/example/slice.go
 */
package main

import "fmt"

func main() {

	//1. 普通初始化一个切片数组
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s: ", s)

	//2. 普通初始化一个哈希表，并且查询这个值有没有
	mp := make(map[int]int)
	mp[1] = 2
	mp[2] = 3
	if val, ok := mp[3]; ok {
		fmt.Println(val)
	}

	//3. 普通切片数组追加元素.追加后如果超过元素本身的cap，那么返回的是一个新的slice,元素底层地址完全变了
	//输出：
	//s = 6   0xc00005c190
	//s = 6   0xc00005c190
	//s = 6   0xc00005c190
	//s = 12   0xc000102010
	s = append(s, "d")
	fmt.Printf("s = %d   %p\n", cap(s), &s[1])
	s = append(s, "d")
	fmt.Printf("s = %d   %p\n", cap(s), &s[1])
	s = append(s, "d")
	fmt.Printf("s = %d   %p\n", cap(s), &s[1])
	s = append(s, "d")
	fmt.Printf("s = %d   %p\n", cap(s), &s[1])

	//4. len同样可以得到切片数组的长度
	fmt.Println(len(s))

	//5.数组切片这样是将s切片中的0 - 1个元素浅浅浅浅拷贝到l中
	l := s[0:2]
	fmt.Println("l = ", l)
	fmt.Printf("l = %d   %p\n", cap(l), &l[1])

	//6. 切片深拷贝, 同样支持[:]
	m := make([]string, len(s))
	copy(m, s[:])
	m[1] = "5"
	fmt.Println("m, s = ", m, s)

	//7. 多维切片数组,每个内部切片数组长度可以不一样，类似vector<vector<int>>
	// 同样可以追加
	k := make([][]int, 3)

	for i := 0; i < len(k); i++ {
		innerlen := i + 1
		k[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			k[i][j] = i + j
		}
	}
	k = append(k, []int{1, 2, 3, 4})
	fmt.Println("k = ", k)
}
