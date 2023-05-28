package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(1 * time.Second)

	<-timer1.C
	fmt.Println("afer 1 second")

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("after 2 second")
	}()

	stop := timer2.Stop()

	if stop {
		fmt.Println("timer2 is stop")
	}

	var i = 0
	done := make(chan bool)
	go func(num *int) {
		for i := 0; i < 10; i++ {
			*num++
		}
		done <- true
	}(&i)

	<-done
	fmt.Println(i)

}
