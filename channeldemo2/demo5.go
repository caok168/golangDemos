package main

import (
	"fmt"
	"time"
)

func worker5_1(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func worker5_2(id int, c chan int) {
	for n := range c {
		fmt.Printf("Workder %d received %d\n", id, n)
	}
}

func channelClose() {
	c := make(chan int)

	go worker5_2(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)

}

func main() {
	channelClose()
}
