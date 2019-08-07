package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().AddDate(0, 0, 6).Weekday())
	fmt.Println(time.Now().AddDate(0, 0, 5).Weekday())
	fmt.Println(time.Now().AddDate(0, 0, 4).Weekday())
	fmt.Println(time.Now().AddDate(0, 0, 3).Weekday())
	fmt.Println(time.Now().AddDate(0, 0, 2).Weekday())
	fmt.Println(time.Now().AddDate(0, 0, 1).Weekday())
	fmt.Println(time.Now().AddDate(0, 0, 0).Weekday())

}
