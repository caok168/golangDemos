package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("test")
}

func Add(x, y int) int {

	res := 0

	res = CalcTriangele(x, y)

	res = CalcTriangele(res, res)

	for i := 0; i < 100; i++ {
		res = res * res
	}

	res = res * res

	return 0
}

func CalcTriangele(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
