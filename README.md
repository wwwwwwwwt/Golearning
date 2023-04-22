<!--
 * @Author: zzzzztw
 * @Date: 2023-03-31 10:16:52
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-22 17:07:33
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

```
go env -w GO111MODULE=on
```

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

## class 6. 导入包的技巧

```go

package main

import (
	_ "class5/lib1" // 匿名导入包，表示不使用这个包的方法，只让这个包内的init（）完成初始化
	. "class5/lib2"  // 表示将lib2 内的作用域完全导入本函数作用域，可以在本作用域直接使用lib2中的函数方法， 但不推荐可能会有函数重名
	aa "fmt" // 给包起别名。 fmt.方法 可以变为 aa.方法
)

func main() {

	//lib1.Lib1test()
	Lib2test()
	aa.Println("hello")
}


```


## class 7. go中的指针与按指针传参

* 基本和c++一样,但基本类型没有按引用传递参数

```go

package main

import (
	"fmt"
)

func changeval(a *int) {
	*a = 20
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {

	var a int = 10
	fmt.Println("a = ", a, "&a = ", &a)

	changeval(&a)

	fmt.Println("a = ", a, "&a = ", &a)

	fmt.Println("-------swap val --------")

	x, y := 100, 200

	fmt.Println("x = ", x, "y = ", y)

	swap(&x, &y)

	fmt.Println("x = ", x, "y = ", y)
}

输出：
a =  10 &a =  0xc0000ba000
a =  20 &a =  0xc0000ba000
-------swap val --------
x =  100 y =  200
x =  200 y =  100

```

## class 8. defer的调用机制

* defer + 函数或语句，在函数完全结束时（在return之后）触发
* 多个defer 一起使用时类似压栈，先声明的defer先被压进栈，后触发

```go

package main

import (
	"fmt"
)

func foo(a int) int {
	fmt.Println("main end ", a)
	return 0
}

func foo2() int {
	defer foo(1)
	return foo(2)
}

func main() {
	/*defer foo(1)
	defer foo(2)
	defer foo(3) */

	foo2()

}

输出：
main end  2
main end  1

```

## class 9. 固定数组与动态数组

1. 固定数组

```go

var arr1 [10] int // 十个默认值0

arr2 := [10]int{1,2,3,4} // 1234 + 6个0



func foo(a [10]int)
注意这种数组传参为按值传递

func foo( a[]int)为按引用传递

使用时

foo(arr2[起始位置:终止位置])

[起始位置，终止位置) 左闭右开不包含终止位置

表示将起始位置到终止位置之前的那个位置切片成一个数组传入函数

```

2. 动态数组

```go

a := []int{1,2,3,4}

```

## class10. 声明slice数组的四种方法和判断slice是否为空

```go

package main

import (
	"fmt"
)

func main() {

	//方法1 声明长度为3的切片数组

	slice1 := []int{1, 2, 3}

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 方法2 声明长度为0的切片数组,没有分配容量，可以通过make([]类型，长度)进行分配空间，值为默认值

	var slice2 []int

	slice2 = make([]int, 4)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	//方法3， 通过var 声明

	var slice3 []int = make([]int, 5)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	// 方法4， 动态直接初始化，分配空间

	slice4 := make([]int, 6)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	// 判断是否为空，就是判断是否为nil
	var a []int

	if a == nil {
		fmt.Println("nil")
	} else { // else 要和右括号在一行
		fmt.Println("not nil")
	}
}

```

## class 11. slice的切片追加和截取

* 类似vector底层动态扩容的过程，底层有size 和 cap 两个成员，size超过了cap的大小 cap容量就会翻倍。

* 使用切片名 = append（切片名，值），将这个值追加进切片数组，类似vector中的push_back
```go

package main

import (
	"fmt"
)

func main() {

	a := make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(a), cap(a), a)

	a = append(a, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(a), cap(a), a)

	a = append(a, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(a), cap(a), a)

	a = append(a, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(a), cap(a), a)
}


```

2. 切片的截取

* 使用时 s1 = arr2[起始位置:终止位置]
* [起始位置，终止位置) 左闭右开不包含终止位置
* 注意，此时s1指向的地址就是arr2的地址，s1的修改会影响arr2
* 如果想完全分开，就需要分配一个空slice数组s2, 然后使用copy(s2, arr2[:])完全复制过去

```go

s2 := a[3:6]
s2[0] = 100

输出：
len = 3, cap = 7, slice = [100 2 2]
len = 6, cap = 10, slice = [0 0 0 100 2 2]
---------------
s2 := make([]int, 3)

copy(s2, a[3:6])

s2[0] = 100

输出：
len = 3, cap = 3, slice = [100 2 2]
len = 6, cap = 10, slice = [0 0 0 2 2 2]
这样修改s2不会影响到a数组
```


## class12. map的三种定义方式

```go

	//定义mp 为map[key] val类型的map,但这样声明没有底层空间，是nil
	var mp map[string]string
	if mp == nil {
		fmt.Println("map is nil")
	}

	// 方法1：指定底层cap大小
	mp = make(map[string]string, 2) //给mp分配2个空间， 一旦满了，就会自动动态扩容
	mp["abc"] = "123"
	mp["c"] = "123"
	mp["d"] = "123"
	mp["e"] = "123"
	fmt.Printf("size = %d, map = %v\n", len(mp), mp)


	// 方法2：不指定大小，会自动扩容
	mp2 := make(map[int]string)
	mp2[1] = "abc"
	mp2[3] = "bcd"
	fmt.Printf("map = %v\n", mp2)


	// 方法3， 直接以值初始化
	mp3 := map[int]string{
		1: "abc",
		2: "cdf",
		3: "rdf",
	}

	fmt.Printf("map = %v\n", mp3)


```