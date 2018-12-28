package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	var x = 10
	var y = 20
	a, b := swap(x, y)
	fmt.Println("Initial order: ", x, y)
	fmt.Println("After swapping: ", a, b)
}
