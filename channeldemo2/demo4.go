package main

func bufferedChannel() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
}

func main() {
	bufferedChannel()
}
