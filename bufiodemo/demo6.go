package main

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadRune 从 b 中读出一个 UTF8 编码的字符并返回
// 同时返回该字符的 UTF8 编码长度
// 如果 UTF8 序列无法解码出一个正确的 Unicode 字符
// 则只读出 b 中的一个字节，并返回 U+FFFD 字符，size 返回 1
//func (b *Reader) ReadRune() (r rune, size int, err error)

// UnreadRune 撤消最后一次读出的 Unicode 字符
// 如果最后一次执行的不是 ReadRune 操作，则返回一个错误
// 因此，UnreadRune 比 UnreadByte 更严格
//func (b *Reader) UnreadRune() error

func main() {
	s := strings.NewReader("你好，世界！")
	br := bufio.NewReader(s)

	c, size, _ := br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	br.UnreadRune()
	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)
}
