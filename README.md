<!--
 * @Author: zzzzztw
 * @Date: 2023-03-31 10:16:52
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 13:25:32
 * @FilePath: /Golearning/README.md
-->
# Golearning

## class1
1. go run + 文件名 = 编译加运行   
2. go build + 文件名 + ./文件名 编译加运行
3. 主函数main所在文件要有 package main 叫main包
4. import （包名）导入多个包
5. go语言要求{}的左括号与函数名在同一行
6. 语句加不加分号无所谓，一般推荐不加
```go
package main // 程序的包名

import (
	"fmt"
	"time"
)

func main() {

	//语句加不加分号无所谓，一般推荐不加
	//go语言要求{}的左括号与函数名在同一行，如上所写
	fmt.Println("hello world")
	time.Sleep(1 * time.Second)
}


```

## class2. 声明变量

```go

// 方法1， 声明一个变量, 没给初始化值为0

var a int

var b int = 10

// 方法2，声明一个变量，给一个初始化值

var b int  = 100

//方法3， 省去数据类型，靠值自动推断

var c = 100

//方法4， 省去var关键字, 比较常用，但不能声明全局变量,方法123可以用作声明全局变量

d := "abcd"

// 多变量初始化
xx, yy := 100, "abc"
var zz, ee = 100, "abc"

```

## class3. const 和 iota

1. const 用法就是把var 的位置换成const

```go

const a int = 8

```

2. 通过const 定义枚举类型

```go

const (
	BEIJING = 0
	SHANGHAI = 1
	SHENZHEN = 2
)

```

3. 通过向const内添加关键字iota，使得每行的iota都会默认+1， 第一行默认为0

* iota 只能用到const内部

```go

const (

	BEIJING = iota // 0
	SHANGHAI // 1
	SHENZHEN //2

)


const (

	BEIJING = 10 * iota
	SHANGHAI  // 10
	SHENZHEN  // 20
)


const (

	a, b = iota + 1, iota + 2 // 1, 2
	c, d // 2, 3
	e, f // 4, 5

	g, h = iota * 2, iota * 3 // 6,9
	i, k // , 12

)

```

## class4. 函数多返回值


```go

// func 函数名（参数 参数类型） 返回值类型 匿名
func foo1(a string, b int) (int, int) {

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	d := 200
	return c, d

}


// func 函数名（参数 参数类型） 返回值类型 有名
func foo2(a string, b int) (c int, d int) {

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// c d 的作用域空间时foo2 的整个空间
	fmt.Println("c = ", c) // 默认为0
	fmt.Println("d = ", d) // 默认为0
	c = 100
	d = 200
	return c, d

}


package main


import "fmt"


func swap(x, y string) (string, string) {
   return y, x
}


func main() {
   a, b := swap("Mahesh", "Kumar")
   fmt.Println(a, b)
}

//以上实例执行结果为：

//Kumar Mahesh


```

## class5. init函数和import导包路径

* 使用go mod管理包
* 首先我们需要初始化一个go.mod，使用

```go
go mod init test
# test可以是任意的名称

```

* 然后我们引入模块的时候，以test(初始化时定义)开头，然后接模块路径，比如

```go
import "test/module"
```
* 如果我们想要引用嵌套模块也是一样的

```go
import "test/module/moduleA"
```

1. 进入包后首先会找这个包依赖的包，依次递归到最深的包
2. 依次找到它的const 全局变量等，最后找到它的init函数，执行
3. 包内的对外方法首字母一定要大写开头，小写开头对应着私有方法对外不可见。