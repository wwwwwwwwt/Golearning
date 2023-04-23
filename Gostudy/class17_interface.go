/*
 * @Author: zzzzztw
 * @Date: 2023-04-24 00:29:20
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-24 00:37:13
 * @FilePath: /Golearning/Gostudy/class17_interface.go
 */
package main

import "fmt"

type Animal interface {
	Sleep()
	Getcolor() string
	Gettype() string
}

type Cat struct { // 继承基类接口animal 不需要写上，只需要定义animal中所有的函数即可
	color string
	kind  string
}

func (t *Cat) Sleep() {
	fmt.Println("Cat is sleeping.....")
}

func (t *Cat) Getcolor() string {
	return t.color
}

func (t *Cat) Gettype() string {
	return t.kind
}

type Dog struct {
	color string
	kind  string
}

func (t *Dog) Sleep() {
	fmt.Println("Dog is sleeping.....")
}

func (t *Dog) Getcolor() string {
	return t.color
}

func (t *Dog) Gettype() string {
	return t.kind
}

func main() {
	var p Animal

	cat := Cat{"black", "cat"}

	dog := Dog{"yellow", "dog"}

	p = &cat
	p.Getcolor()
	p.Gettype()
	p.Sleep()

	fmt.Println("-------------------")

	p = &dog
	p.Getcolor()
	p.Gettype()
	p.Sleep()
}
