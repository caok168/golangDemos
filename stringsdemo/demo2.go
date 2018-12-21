package main

import (
	"fmt"
	"strings"
)

func main() {
	//isok := KeyMatch("/foo/bar", "/foo/*")
	//fmt.Println(isok)

	//splitNDemo()

	//stringCountDemo()

	str := "1231231231 # this is a comment"
	res := removeComment(str)
	fmt.Println(res)
}

func trimPrefixTest() {
	auth := "Bearer 12312313"
	if len(auth) > 0 {
		token := strings.TrimPrefix(auth, "Bearer ")

		fmt.Printf("token:%s", token)
	}
}

// For example, "/foo/bar" matches "/foo/*"
func KeyMatch(key1 string, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	if len(key1) > i {
		return key1[:i] == key2[:i]
	}
	return key1 == key2[:i]
}

func splitNDemo() {
	p := strings.SplitN("users:create", ":", 2)
	fmt.Println(p[0])
	fmt.Println(p[1])
}

func stringCountDemo() {
	i := strings.Count("aaaaabbbbaaaa", "444")
	fmt.Println("count:", i)
}

func removeComment(s string) string {
	pos := strings.Index(s, "#")
	if pos == -1 {
		return s
	}
	return strings.TrimSpace(s[0:pos])
}
