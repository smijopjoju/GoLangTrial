package main

import "fmt"

func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	adder1 := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(adder1(i))
	}
}
