package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	date := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	datetime, _ := time.Parse("2006-01-02", date)

	year, month, day := time.Now().Date()

	fmt.Println(strconv.Itoa(year) + "-" + month.String() + "-" + strconv.Itoa(day))

	fmt.Println(date)

	fmt.Println(datetime)

	aaa := datetime.Format("2006-01-02")
	fmt.Println(aaa)
}
