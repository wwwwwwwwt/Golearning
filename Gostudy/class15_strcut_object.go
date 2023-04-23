/*
 * @Author: zzzzztw
 * @Date: 2023-04-23 09:48:24
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-04-24 00:16:53
 * @FilePath: /Golearning/Gostudy/class15_strcut_object.go
 */
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
