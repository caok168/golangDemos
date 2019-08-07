package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Contract struct {
	empId    int
	basicpay int
}

type Permanent struct {
	empId    int
	basicpay int
	jj       int
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.jj
}

func main() {
	pemp1 := Permanent{empId: 1, basicpay: 100, jj: 10}

	salary := pemp1.CalculateSalary()

	fmt.Println(salary)

	cemp1 := Contract{empId: 2, basicpay: 1000}
	salary2 := cemp1.CalculateSalary()
	fmt.Println(salary2)
}
