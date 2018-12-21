package main

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadLine 是一个低级的原始的行读取操作
// 大多数情况下，应该使用 ReadBytes('\n') 或 ReadString('\n')
// 或者使用一个 Scanner
//
// ReadLine 通过调用 ReadSlice 方法实现，返回的也是缓存的切片
// ReadLine 尝试返回一个单行数据，不包括行尾标记（\n 或 \r\n）
// 如果在缓存中找不到行尾标记，则设置 isPrefix 为 true，表示查找未完成
// 同时读出缓存中的数据并作为切片返回
// 只有在当前缓存中找到行尾标记，才将 isPrefix 设置为 false，表示查找完成
// 可以多次调用 ReadLine 来读出一行
// 返回的数据在下一次读取操作之前是有效的
// 如果 ReadLine 无法获取任何数据，则返回一个错误信息（通常是 io.EOF）
//func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

func main() {
	//s := strings.NewReader("ABC\nDEF\r\nGHI\r\nJKL")

	s := strings.NewReader("ABCDEF\r\nGHI\r\nJKL\n111")
	br := bufio.NewReader(s)

	w, isPrefix, _ := br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
}
