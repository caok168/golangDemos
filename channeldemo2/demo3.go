package main

import (
	"fmt"
	"time"
)

// 发送数据的channel

func createWorker1(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()

	return c
}

func chanDemo2() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker1(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo2()
}
