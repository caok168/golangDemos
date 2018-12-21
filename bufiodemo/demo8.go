package main

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadSlice 在 b 中查找 delim 并返回 delim 及其之前的所有数据的切片
// 该操作会读出数据，返回的切片是已读出数据的引用
// 切片中的数据在下一次读取操作之前是有效的
//
// 如果 ReadSlice 在找到 delim 之前遇到错误
// 则读出缓存中的所有数据并返回，同时返回遇到的错误（通常是 io.EOF）
// 如果在整个缓存中都找不到 delim，则 err 返回 ErrBufferFull
// 如果 ReadSlice 能找到 delim，则 err 始终返回 nil
//
// 因为返回的切片中的数据有可能被下一次读写操作修改
// 因此大多数操作应该使用 ReadBytes 或 ReadString，它们返回的不是数据引用
//func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

func main() {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadSlice(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(' ')
	fmt.Printf("%q\n", w)
}
