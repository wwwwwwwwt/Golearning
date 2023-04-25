<!--
 * @Author: zzzzztw
 * @Date: 2023-03-31 10:16:52
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-25 13:38:41
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


## class 12. map的三种定义方式

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

## class 13. map的基本使用方式，增删改查


```go

package main

func changemap(a map[string]string) { // 注意mao传参是按引用传参，动态slice也是
	a["italy"] = "bcd"
}

func main() {

	citymap := make(map[string]string)

	// 添加元素
	citymap["China"] = "beijing"
	citymap["USA"] = "NK"
	citymap["UK"] = "london"

	//查询这个元素在不在map中
	if val, ok := citymap["JAPAN"]; ok {
		println("Japan city is ", val)
	} else {
		println("japan is nil")
	}

	if val, ok := citymap["China"]; ok != false {
		println(val, "is the city of China")
	} else {
		println("nil")
	}

	// 遍历元素
	for key, val := range citymap {
		println("country is ", key, "city is", val)
	}

	println("-----------------------")

	//删除键
	delete(citymap, "UK")
	for key, val := range citymap {
		println("country is ", key, "city is", val)
	}

	//修改键
	citymap["USA"] = "abc"

	//按引用传参
	changemap(citymap)
	for key, val := range citymap {
		println("country is ", key, "city is", val)
	}

}

```
]

## class 14. 结构体struct的基本用法


```go

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


```
## class 15. Go语言的class是基于struct 和 struct 对象的方法来实现的

* 特别注意：方法的首字母大写是对包外可见的方法类似public，小写的方法只在包内可见类似private

* 类的方法需要传入类对象的指针，否则传入的是副本，不能修改类的成员变量

```go
package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (t *Hero) GetName() string { // t 为类对象别名
	return t.Name
}
func (t *Hero) Show() {
	fmt.Printf("name: %s ad: %d Level: %d\n", t.Name, t.Ad, t.Level)
}

func (t *Hero) SetName(newname string) { // 修改类成员需要传入对象指针，不传指针的话，是传入该方法对象的拷贝
	t.Name = newname
}

func main() {

	hero := Hero{"ztw", 99999, 99999}
	hero2 := Hero{Name: "wgl", Ad: 1000, Level: 1000}

	hero.Show()

	hero.SetName("zzzztw")

	hero.Show()

	hero2.Show()
}


```


## class 16. 类的继承1，这种继承实现不了多态

```go

package main

import "fmt"

type Human struct { // 定义父类
	name string
	sex  string
}

type Superman struct { // 定义子类，直接把父类放在子类
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



```

## class 17. 利用interface实现多态，类似抽象类

* interface本身是一个指针
* interface接口类似抽象类，子类需要实现interface内定义的所有函数，就相当于继承
* 使用时将interface绑定到子类对象的地址上，就可以多态的使用函数了


```go
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
	//使用cat的多态
	p = &cat
	p.Getcolor()
	p.Gettype()
	p.Sleep()

	fmt.Println("-------------------")
	//使用dog的多态
	p = &dog
	p.Getcolor()
	p.Gettype()
	p.Sleep()
}


```


## class 18. 万能类型空接口 interface{}类似多态
* 类似泛型
* 1.18version以后可以使用any 代替interface{}
* int, string, float32, float64, struct内部都实现了interface{}接口
* 就可以用interface{}类型，引用任意的数据类型。
* 类型断言机制，利用if 或 switch进行判断， 注意interface{}返回类型需要&

```go

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

	c := myfunc(b).(*a) // 后面是断言赋值可以用来强制转换

	fmt.Printf("type is %T, c is %s\n", c, c.name)

}



```

## class 19. 变量的内置类型pair结构

* 变量中包括：type,val
* type分为static type 和 dynamic type, 动态类型指的是interface所指向的类型

```go

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


```

## class 20. reflect反射机制用法

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	// 获取类型和值，reflect.TypeOf(), reflect.ValueOf()
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("val : ", reflect.ValueOf(arg))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (t User) Call() {
	fmt.Println("user is call...")
	fmt.Printf("%v\n", t)
}

func DofieldandMethod(input interface{}) {

	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is", inputType)

	//获取input的val
	inputval := reflect.ValueOf(input)
	fmt.Println("input val is", inputval)

	//通过type获取里面的字段
	//1. 获取interface{}中的reflect type类型中的.NumField， 进行遍历
	//2. 得到每个type类型的Field(i)，是数据类型
	//3. 通过每个val类型中的Filed(i)中的Interface()接口得到val

	for i := 0; i < inputType.NumField(); i++ { // 注意，获取字段时需要结构体内的变量首字母大写，否则不能获取到类似private
		field := inputType.Field(i)
		val := inputval.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, val)
	}

	// 通过type获得里面的方法并调用

	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)

		fmt.Printf("%s: %v\n", method.Name, method.Type)

	}

}

func main() {

	var num float64 = 1.2345
	reflectNum(num)

	user := User{1, "ztw", 24}
	DofieldandMethod(user)

}

输出：

type:  float64
val :  1.2345
input type is main.User
input val is {1 ztw 24}
Id: int = 1
Name: string = ztw
Age: int = 24
Call: func(main.User)
```


## class 21. go中通过反射解析结构体标签Tag

* 传入对象指针就用 reflect.TypeOf().Elem()
* 传入对象本身就用 reflect.TypeOf()
* 通过Field().  获得一系列参数

```go

package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"` // 标签类似注释，可以通过反射得到
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() // 传入指针用TypeOf().Elem

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", taginfo, "doc:", tagdoc)
	}
}

func main() {

	var re resume
	findTag(&re)
}


```


## class 22. 结构体标签再json中的应用

* json包 ```"encoding/json"```
* 序列化 Marshal struct -->json 反序列化 Unmarshal json ---> struct, 反序列化时，传入struct对象的指针

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name  string   `json:"title"`
	Year  int      `json:"year"`
	Price int      `json:"rmb"`
	Actor []string `json:"actors"`
}

func main() {

	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

	// 编码的过程， struct ---> json
	jsonstr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal is error")
		return
	} else {
		fmt.Printf("jsonstr = %s\n", jsonstr)
	}

	// 解码的过程， json ---> struct

	myMovie := Movie{}                      // 建立结构体接收
	err = json.Unmarshal(jsonstr, &myMovie) // 以指针传入

	fmt.Printf("%v\n", myMovie)
}



```


##  class 23. 创建goroutine

* 通过go 关键字

```go

package main

import (
	"fmt"
	"runtime"
	"time"
)

func foo() {
	fmt.Println("hhhhhhh")
}

func main() {

	// 匿名函数执行
	go func() {
		defer fmt.Println("A defer")

		func() {
			defer fmt.Println("B defer")
			runtime.Goexit() //退出当前协程，直接退出，到执行defer那一步 / return是退出当前函数，还会执行上层函数体的后续
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	// 有名函数执行
	go foo()

	for {
		time.Sleep(1 * time.Second)
	}
}

```


## class 24. channel的基本定义与使用，channel用于两个routine的数据传递管道

* channel 有同步两个线程的作用，类似c++ promise
* c := make(chan Type)建立channel，type指定传递类型，无缓冲,无缓冲的channel当数据被放进管道时，数据没被拿走，这个channel会被锁住，直到数据被拿走
* c := make(chan Type, capicity) 建立有缓冲的channel
* 子线程 c <- 传递的值
* 主线程 num := <-c传递给主线程值
* 类似管道，一边进一边出

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine is finish") // defer 会在c <-666之后执行
		c <- 666                                 // 在此处就把值通过管道传递给了主线程num，在此处同步，然后接下来继续执行下面的
		time.Sleep(1 * time.Second)
	}()

	if num := <-c; num != 0 {
		fmt.Println("num is ", num)
	}
	//fmt.Println("num is", num)

	time.Sleep(2 * time.Second)
}

```

## class 25. 有缓存的channel

* 有缓冲区的channel当channel内元素小于缓冲区cap大小时，不会发生阻塞，当元素达到上限时会阻塞

```go
定义
c := make(chan int, 缓冲区大小)
```

* 示例
```go

package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)

	fmt.Println("len c = ", len(c), "cap c =", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		c <- 666
		c <- 777
		c <- 888
		c <- 999 // 不写这一行无论主线程有没有取出元素，都会直接结束，写了这行主线程不取元素就会阻塞
	}()

	time.Sleep(1 * time.Second)

	num1 := <-c
	num2 := <-c
	num3 := <-c

	fmt.Println("1 :", num1, "2 :", num2, "3 :", num3)
	time.Sleep(1 * time.Second)
}
输出
len c =  0 cap c = 3
1 : 666 2 : 777 3 : 888
子go程结束

```


## class 26. 使用close关闭管道

* 只有当你确实没有任何发送数据了，或者向现实的结束range循环时，一般才关闭channel
* 向关闭的channel再次发送数据会报错
* 关闭channel后还可以继续从channel缓存中接受数据
* 当向nil channel（不使用make声明的channel变量）无论收发数据都会被阻塞，所以都应该使用make声明channel

```go
package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			c <- i
		}

		close(c) // close 关闭一个管道，不写这行运行会报死锁的错误

	}()

	for {

		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main finished")

}

```


## class 27. channel和range

* for date := range c 阻塞取出channel内数据，直到channel被关闭

```go

package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}
		close(c)
	}()

	for data := range c { // 只要channel中有数据，就阻塞尝试取出，直到channel被关闭。channel一直不关闭会导致死锁
		fmt.Println(data)
	}

}

```

## class 28. channel 和 select

* select 类似c++中的select io多路复用，同时监控多个channel的可读可写,用法类似switch case
* 注意逻辑，防止线程死锁

```go
package main

import "fmt"

func main() {

	c := make(chan int)
	quit := make(chan int)
	go func() {
		x, y := 1, 1

		for {
			select {
			case c <- x:
				//x, y = y, x+y // 斐波那契数列
				z := x
				x = y
				y = z + y
			case <-quit:
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		data := <-c

		fmt.Println(data)
	}
	quit <- 0
}




```

## class 29. go modules工作模式

导入包后
go mod init + 名字或github地址
go mod tidy