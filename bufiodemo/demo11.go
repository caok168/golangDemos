package main

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadString 功能同 ReadBytes，只不过返回的是一个字符串
//func (b *Reader) ReadString(delim byte) (line string, err error)

func main() {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadString(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)
}
