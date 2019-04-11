package main

import "fmt"

func main() {

	nums := []int{1, 2}

	change(nums)

	for _, v := range nums {
		fmt.Println(v)
	}

}

func change(nums []int) {
	nums = append(nums, 3)

	for _, v := range nums {
		fmt.Println(v)
	}
}
