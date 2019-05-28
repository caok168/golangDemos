package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("test")

	p, err := plugin.Open("./main.so")
	if err != nil {
		panic(err)
	}

	m, err := p.Lookup("GetProductBasePrice")
	if err != nil {
		panic(err)
	}

	res := m.(func(int) int)(30)
	fmt.Println(res)
}
