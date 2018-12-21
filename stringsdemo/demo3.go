package main

import "fmt"

func main() {

	str := "abcdefg"

	strNew := str[1 : len(str)+1]

	fmt.Println(strNew)

}
