/*
 * @Author: zzzzztw
 * @Date: 2023-05-26 14:14:28
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-26 14:22:36
 * @FilePath: /Golearning/example/basego/interface.go
 */
package main

import "fmt"

type ge interface {
	area() float64
	perim() float64
}

type circle struct {
	r float64
}

func (c *circle) area() float64 {
	return 3.14 * c.r * c.r
}

func (c *circle) perim() float64 {
	return 3.14 * c.r * 2
}

type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g ge) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {

	r := &rect{width: 2, height: 4} // 注意，传入参数需要与这个这个对象的方法一致。例如，area（）方法是按指针传递对象，那么这里也需要传入指针。
	c := &circle{r: 2}

	measure(r)
	measure(c)
}
