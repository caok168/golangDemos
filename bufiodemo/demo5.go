package main

import (
	"fmt"
	"os"
)

const path = "./volumes/uploads"

func main() {
	err := os.Mkdir(path, 0700)

	if err != nil {
		fmt.Println(err)
	}
}
