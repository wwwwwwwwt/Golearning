/*
 * @Author: zzzzztw
 * @Date: 2023-04-24 18:23:03
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-24 19:09:14
 * @FilePath: /Golearning/Gostudy/class21_reflect_Tag.go
 */
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
