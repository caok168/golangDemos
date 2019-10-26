package main

import "fmt"

type Pet3 interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog1 struct {
	name string // 名字
}

func (dog *Dog1) SetName(name string) {
	dog.name = name
}

func (dog Dog1) Name() string {
	return dog.name
}

func (dog Dog1) Category() string {
	return "dog"
}

func main() {
	var dog31 *Dog1
	fmt.Println("The first dog is nil.[wrap1]")
	dog32 := dog31
	fmt.Println("The second dog is nil.[wrap1]")
	var pet Pet3 = dog32
	if pet == nil {
		fmt.Println("The pet is nil.[wrap1]")
	} else {
		fmt.Println("The pet is not nil.[wrap1]")
		fmt.Println(pet)
	}
}
