package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好，世界")
}

var Greeter greeting
