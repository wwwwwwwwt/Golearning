<!--
 * @Author: zzzzztw
 * @Date: 2023-03-31 10:16:52
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-21 00:22:51
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

## class2 声明变量

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
