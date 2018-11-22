package main

import "fmt"

func main() {
	stu := make(map[string]string)

	stu["name"] = "ck"
	stu["age"] = "29"

	if name, ok := stu["name"]; ok {
		fmt.Println(name)
	}

	if addr, ok := stu["addr"]; ok {
		fmt.Println(addr)
	} else {
		fmt.Println("No addr")
	}
}
