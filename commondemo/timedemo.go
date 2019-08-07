package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()

	fmt.Println(time.Date())

	fmt.Println(time.Format("2006-01-02 15:04:05"))
}
