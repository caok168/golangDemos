package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Read 从 b 中读出数据到 p 中，返回读出的字节数
// 如果 p 的大小 >= 缓存的总大小，而且缓存不为空
// 则只能读出缓存中的数据，不会从底层 io.Reader 中提取数据
// 如果 p 的大小 >= 缓存的总大小，而且缓存为空
// 则直接从底层 io.Reader 向 p 中读出数据，不经过缓存
// 只有当 b 中无可读数据时，才返回 (0, io.EOF)
func main() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	b := make([]byte, 20)

	n, err := br.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)
	//fmt.Printf("%s %v %v\n", b[:n], n, err)

	n, err = br.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)

	n, err = br.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)
}
