package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}

	for i, v := range slice {
		fmt.Println("Index:", i, "Value:", v)
	}
}
