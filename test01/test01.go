package main

import "fmt"

func main() {
	fmt.Println("1111111111111")
	fmt.Println("1111111111111")
	fmt.Println("1111111111111")
	a := 5
	b := 3
	sum := a + b
	if sum > 5 {
		fmt.Println("a+b的和大于5")
	} else {
		fmt.Println("a+b的和小于5")
	}
	fmt.Println("a+b:", a+b)
	chengji(a, b)
}

func chengji(a, b int) int {
	fmt.Println("a*b", a*b)
	return a * b
}
