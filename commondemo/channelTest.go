package main

import (
	"fmt"
	"time"
)

var g = G{}

func main() {

	fmt.Println(g)

	for i := 0; i < 5; i++ {
		go println(i)
		time.Sleep(time.Millisecond)
	}
}

func PrintTest(i int) {
	println(i)
}

type G struct {
}
