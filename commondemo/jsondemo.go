package main

import (
	"fmt"
	"strings"
)

func main() {

	aaa := strings.TrimRight("uploads/a4a5c64e1d3c6d83de6cc123966f2888.bin", ".bin")
	fmt.Println(aaa)

	//strings.TrimPrefix("hank_goGuide", "hank_")

	bbb := strings.TrimPrefix("uploads/a4a5c64e1d3c6d83de6cc123966f2888", "uploads/")
	fmt.Println(bbb)

	fileID := strings.TrimLeft(strings.TrimRight("uploads/a4a5c64e1d3c6d83de6cc123966f2888.bin", ".bin"), "uploads")

	fmt.Println(fileID)

	//person := &Person{
	//	Id:   1,
	//	Name: "ck",
	//}
	//
	//byte, err := json.MarshalIndent(person, "", "	")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//content := string(byte)
	//
	//fmt.Println(content)
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
