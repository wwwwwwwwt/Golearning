/*
 * @Author: zzzzztw
 * @Date: 2023-04-24 16:29:50
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-24 19:06:08
 * @FilePath: /Golearning/Gostudy/class20_reflect.go
 */
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
