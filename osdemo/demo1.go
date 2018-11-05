package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getegid())
	fmt.Println(os.Geteuid())
}
