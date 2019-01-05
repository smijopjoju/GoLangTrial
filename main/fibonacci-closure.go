package main

import (
	"fmt"
	"strconv"
)

func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		sum := first + second
		first = second
		second = sum
		return sum
	}
}

func main() {
	fibRefe := fibonacci()
	var fibSeries string = "1 "
	for i := 0; i < 10; i++ {
		fibSeries += strconv.Itoa(fibRefe()) + " "
	}

	fmt.Println(fibSeries)
}
