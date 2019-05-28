package main

import (
	"bytes"
	"fmt"
)

func main() {
	//fmt.Println("test")
	//
	//i := 10
	//name := "ck"

	str := "hello"
	arr := []byte(str)

	strAll := ""

	for i := 0; i < 5; i++ {
		strAll += str
	}

	strAllByte := []byte(strAll)
	fmt.Println("strAllByte:=", strAllByte)

	eigenArr := make([][]byte, 5)

	var eigens []byte

	fmt.Println(eigenArr)

	for i := 0; i < 5; i++ {
		eigenArr[i] = arr
	}

	fmt.Println(eigenArr)

	sep := []byte("")
	eigens = bytes.Join(eigenArr, sep)

	fmt.Println(arr)

	fmt.Println(eigens)

}

type Stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
