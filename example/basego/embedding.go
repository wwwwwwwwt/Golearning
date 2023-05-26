/*
 * @Author: zzzzztw
 * @Date: 2023-05-26 14:23:23
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-26 21:59:30
 * @FilePath: /Golearning/example/basego/embedding.go
 */
package main

import "fmt"

type base struct {
	num int
}

func (b *base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func (c *container) say() {
	fmt.Println(c.base.describe(), c.str)
}

func main() {

	co := &container{base: base{10}, str: "some name"}

	fmt.Println(co.describe())
	co.say()
	
}
