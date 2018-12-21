package main

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadBytes 在 b 中查找 delim 并读出 delim 及其之前的所有数据
// 如果 ReadBytes 在找到 delim 之前遇到错误
// 则返回遇到错误之前的所有数据，同时返回遇到的错误（通常是 io.EOF）
// 只有当 ReadBytes 找不到 delim 时，err 才不为 nil
// 对于简单的用途，使用 Scanner 可能更方便
//func (b *Reader) ReadBytes(delim byte) (line []byte, err error)

func main() {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadBytes(' ')
	fmt.Printf("%q\n", w)
	// "ABC "
	w, _ = br.ReadBytes(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadBytes(' ')
	fmt.Printf("%q\n", w)
}
