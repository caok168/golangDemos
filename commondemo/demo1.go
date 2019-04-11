package main

import "fmt"

func main() {
	fmt.Println("test")

	i := 10
	name := "ck"

}

type Stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
