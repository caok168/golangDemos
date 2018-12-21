package main

import (
	"bufio"
	"fmt"
	"strings"
)

//reset丢弃任何的缓存数据，丛植所有状态并且将缓存读切换到r

func main() {
	s := strings.NewReader("ABCDEFG")
	str := strings.NewReader("123455")
	br := bufio.NewReader(s)
	b, _ := br.ReadString('\n')
	fmt.Println(b)
	br.Reset(str)
	b, _ = br.ReadString('\n')
	fmt.Println(b)
}
