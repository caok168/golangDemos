package main

import (
	"time"
	"os"
)

//Channel Timeout 事例
func main() {
	c := make(chan int, 100)

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
			time.Sleep(time.Second)
		}

		os.Exit(0)
	}()

	for {
		select {
		case n := <-c:
			println(n)
		case <-timeAfter(time.Second * 2):
		}
	}
}

func timeAfter(d time.Duration) chan int {
	q := make(chan int, 1)

	time.AfterFunc(d, func() {
		q <- 1
		println("run") 		// 重点在这里
	})

	return q
}
