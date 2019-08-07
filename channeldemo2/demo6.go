package main

import (
	"fmt"
)

func doWorker6(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)

		go func() {
			done <- true
		}()
	}
}

type workerItem struct {
	in   chan int
	done chan bool
}

func createWorker6(id int) workerItem {
	w := workerItem{
		in:   make(chan int),
		done: make(chan bool),
	}

	go doWorker6(id, w.in, w.done)

	return w
}

func chanDemo6() {
	var workers [10]workerItem

	for i := 0; i < 10; i++ {
		workers[i] = createWorker6(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them

	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo6()
}
