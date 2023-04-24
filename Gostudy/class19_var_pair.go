package main

import "fmt"

// 定义接口类型
type Reader interface {
	Readbook()
}
type Writer interface {
	Writebook()
}

// 具体类型,实现了接口类型Reader和Writer的方法，就相当于多重继承
//每个interface接口都可以指向这个子类，并调用自己的方法实现类似多态
type book struct {
}

func (t *book) Readbook() {
	fmt.Println("reading book...")
}

func (t *book) Writebook() {
	fmt.Println("Writing book...")
}

func main() {

	// b: pair<type: book, val:book的地址>
	b := book{}

	//reader: pair<type: 接口类型的type为指向对象的type，这里为book， val：book的地址>
	var reader Reader
	reader = &b
	reader.Readbook()

	//writer: 对另一个接口类型reader进行类型断言
	//因为 book也重新定义了writer的全部函数，所以writer也可以指向book，类似强转
	var writer Writer
	writer = reader.(Writer)
	//	writer = &b
	writer.Writebook()

}
