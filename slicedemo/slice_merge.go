package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3}
	s2 := []int{4, 5, 6, 7}

	s2 = append(s2, s1...)
	fmt.Println(s2)
}
