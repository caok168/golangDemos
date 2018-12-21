package main

import (
	"fmt"
	"time"
)

func main() {

	tm2, _ := time.Parse("01/02/2006", "12/22/2015")

	age := CalcAge(tm2)

	fmt.Println(tm2)

	fmt.Println("age:", age)

	f := float32(int(222))
	fmt.Println(f)

}

// 根据生日计算年龄
func CalcAge(birthday time.Time) int {
	age := time.Now().Year() - birthday.Year()

	if int(time.Now().Month()) < int(birthday.Month()) {
		age--
	}

	return age
}
