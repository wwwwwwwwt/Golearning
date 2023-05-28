package main

import "fmt"

func main() {

	buf := make(chan string, 2)

	buf <- "one"
	buf <- "two"
	close(buf)
	// 非空管道也可以关闭
	for i := range buf {
		fmt.Println(i)
	}

}
