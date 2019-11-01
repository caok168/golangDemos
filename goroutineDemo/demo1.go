package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println()
		}()
	}

	//time.Sleep(time.Millisecond * 500)
	time.Sleep(time.Second)
}
