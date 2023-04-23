/*
 * @Author: zzzzztw
 * @Date: 2023-04-23 08:45:42
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-23 09:01:54
 * @FilePath: /Golearning/Gostudy/class13_usemap.go
 */
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
