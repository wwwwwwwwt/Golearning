/*
 * @Author: zzzzztw
 * @Date: 2023-04-24 00:17:22
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-24 00:27:45
 * @FilePath: /Golearning/Gostudy/class16_inherited.go
 */
package main

import "fmt"

type Human struct {
	name string
	sex  string
}

type Superman struct {
	Human
	level int
}

func (t *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (t *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

func (t *Superman) Eat() {
	fmt.Println("Superman.Eat()...")
}

func (t *Superman) Fly() {
	fmt.Println("Superman.Fly()...")
}

func main() {

	a := Human{name: "ztw", sex: "male"}

	a.Eat()
	a.Walk()

	// 第一种声明子类继承的方法
	b := Superman{Human{"ztw", "male"}, 99999}

	// 第二种，直接.变量再赋值
	/*
		b.name = "ztw"
		b.sex = "male"
		b.level = 99999*/
	b.Eat()  //子类重定义父类的方法
	b.Fly()  // 子类自己的新方法
	b.Walk() // 父类的方法
}
