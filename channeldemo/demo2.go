package main

import (
	"time"
	"os"
)

func main() {
	c := make(chan int, 100)

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
			time.Sleep(time.Second*2)
		}

		os.Exit(0)
	}()

	to := time.NewTimer(time.Second)

	for {
		to.Reset(time.Second)
		select {
		case n := <-c:
			println(n)
		case <-to.C:
			println("超时")
			return
		}
	}
}
