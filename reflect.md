<!--
 * @Author: zzzzztw
 * @Date: 2023-04-30 18:43:13
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-30 18:56:13
 * @FilePath: /zhang/Golearning/reflect.md
-->
# 反射

https://www.jianshu.com/p/b46b1ccd2757

## Golang的反射reflect
### reflect的基本功能TypeOf和ValueOf
```go
// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i.  ValueOf(nil) returns the zero 
func ValueOf(i interface{}) Value {...}

翻译一下：ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0


// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
func TypeOf(i interface{}) Type {...}

翻译一下：TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil

reflect.TypeOf： 直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型
reflect.ValueOf：直接给到了我们想要的具体的值，如1.2345这个具体数值，或者类似&{1 "Allen.Wu" 25} 这样的结构体struct的值
说明反射可以将“接口类型变量”转换为“反射类型对象”，反射类型指的是reflect.Type和reflect.Value这两种


type test struct {
	a int
	b int
}

func main() {
	val := &test{1, 2}
	typ := reflect.TypeOf(val)
	argv := reflect.ValueOf(val)
	fmt.Printf("type is %v, argv is %v\n", typ, argv)
}
输出 type is *main.test, argv is &{1 2}
```