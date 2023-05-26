package main

import "fmt"

type Person struct {
	name string
	age  int
}

func Newperson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func main() {
	a := Newperson("wgl", 24)
	fmt.Println((*a).age) //24
	fmt.Println(*a)       //{wgl 24}
	fmt.Println(a)        //&{wgl 24}
}
