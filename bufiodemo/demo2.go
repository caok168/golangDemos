package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据
// 该操作不会将数据读出，只是引用
// 引用的数据在下一次读取操作之前是有效的
// 如果引用的数据长度小于 n，则返回一个错误信息
// 如果 n 大于缓存的总大小，则返回 ErrBufferFull
// 通过 Peek 的返回值，可以修改缓存中的数据
// 但是不能修改底层 io.Reader 中的数据
func main() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	b, _ := br.Peek(5)
	fmt.Printf("%s\n", b)
	b[0] = 'a'
	b, _ = br.Peek(5)
	fmt.Printf("%s\n", b)
}
