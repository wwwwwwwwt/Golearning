/*
 * @Author: zzzzztw
 * @Date: 2023-05-31 10:32:02
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-31 10:39:41
 * @FilePath: /Golearning/example/channel/worker-pool.go
 */
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "stated job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {

	const numsJob = 5
	jobs := make(chan int, numsJob)
	result := make(chan int, numsJob)
	n := time.Now()
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 1; j <= numsJob; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numsJob; a++ {
		fmt.Println(<-result)
	}
	a := time.Now()

	fmt.Println(a.Unix() - n.Unix())
}
