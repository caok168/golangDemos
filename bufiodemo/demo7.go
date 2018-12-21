package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("你好，世界！")
	br := bufio.NewReader(s)

	fmt.Println(br.Buffered())

	br.Peek(1)
	fmt.Println(br.Buffered())
}
