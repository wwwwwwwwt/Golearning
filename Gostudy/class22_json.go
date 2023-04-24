/*
 * @Author: zzzzztw
 * @Date: 2023-04-24 19:24:53
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-24 19:38:21
 * @FilePath: /Golearning/Gostudy/class22_json.go
 */
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
