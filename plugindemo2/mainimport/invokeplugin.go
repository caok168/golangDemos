package main

import (
	"fmt"
	"os"
	"plugin"
)

func main() {

	p, err := plugin.Open("./pluginhello.so")
	if err != nil {
		fmt.Println("error open plugin: ", err)
		os.Exit(-1)
	}
	s, err := p.Lookup("Hello")
	if err != nil {
		fmt.Println("error lookup Hello: ", err)
		os.Exit(-1)
	}
	if hello, ok := s.(func()); ok {
		hello()
	}
}
