package main

import "sync"

func main() {

	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- 1
			if i%2 == 0 {
				println("A")
			}

		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-ch
			if i%2 == 1 {
				println("B")
			}
		}
		wg.Done()
	}()

	wg.Wait()

}
