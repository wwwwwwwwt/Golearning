package main

import "fmt"

//interface{}是万能数据类型，放在传参中实现泛型

func myfunc(arg interface{}) interface{} {
	fmt.Printf("arg is %T\n", arg)
	fmt.Printf("arg is %v\n", arg)

	// 给interface{}提供类型断言的机制，可以根据不同类型写不同业务,用于参数的传入传出
	if val, ok := arg.(string); ok {
		fmt.Println("arg is string", val)
	} else {
		fmt.Println("arg is not string")
	}

	switch arg.(type) {
	case string:
		return &a{
			name: arg.(string),
		}
	default:
		return "123"
	}
	return nil
}

type a struct {
	name string
}

func main() {

	book := a{"111"}
	myfunc(book)

	myfunc("123")

	b := "5"

	c := myfunc(b).(*a) // 后面是强转类型

	fmt.Printf("type is %T, c is %s\n", c, c.name)

}
