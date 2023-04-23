/*
 * @Author: zzzzztw
 * @Date: 2023-04-23 09:06:02
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-23 09:46:35
 * @FilePath: /Golearning/Gostudy/class14_strcut.go
 */
package main

import "fmt"

type myint int

type book struct {
	name string
	auth string
}

func changebook(a *book) { // 注意结构体 需要传入指针能修改内容
	a.name = "c++"
}

func main() {

	var a myint = 5

	println("a = ", a)

	// 初始化方法1，声明再初始化
	var book1 book
	book1.name = "go"
	book1.auth = "ztw"

	fmt.Printf("book is %v\n", book1)

	changebook(&book1)

	fmt.Printf("book is %v\n", book1)

	//初始化方法2， 直接初始化
	b := book{"go", "ztw"}

	fmt.Printf("b is%v\n", b)
}
