package main

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadByte 从 b 中读出一个字节并返回
// 如果 b 中无可读数据，则返回一个错误
// func (b *Reader) ReadByte() (c byte, err error)

// UnreadByte 撤消最后一次读出的字节
// 只有最后读出的字节可以被撤消
// 无论任何操作，只要有内容被读出，就可以用 UnreadByte 撤消一个字节
func main() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)

	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

	br.UnreadByte()
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	c, err := br.ReadByte()
	fmt.Printf("%c,%v\n", c, err)
}
