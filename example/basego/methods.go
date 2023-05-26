/*
 * @Author: zzzzztw
 * @Date: 2023-05-26 14:06:32
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-26 14:12:36
 * @FilePath: /Golearning/example/basego/methods.go
 */
package main

type rect struct {
	width  int
	height int
}

func (r *rect) change() { // 按指针传入， 传入的对象本身，可以改变其变量
	r.width *= 2
	r.height *= 2
}

func (r rect) change2() { // 不会改变当前对象的变量，是一个拷贝
	r.width *= 2
	r.height *= 2
}

func (r *rect) area() {
	println(r.width, r.height)
}

func main() {
	r := &rect{width: 2, height: 4}
	r.area()
	r.change2()
	r.area()
	r.change()
	r.area()
}
