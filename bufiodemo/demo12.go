package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// WriteTo 实现了 io.WriterTo 接口
//func (b *Reader) WriteTo(w io.Writer) (n int64, err error)

func main() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	b := bytes.NewBuffer(make([]byte, 0))

	br.WriteTo(b)
	fmt.Printf("%s\n", b)
}
