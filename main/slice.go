package main

import "fmt"

func main() {
	numb := [5]int{1, 2, 3, 4, 5}
	var slice []int = numb[0:3]
	var slice1 []int = make([]int, 10, 10)
	fmt.Println(slice)
	fmt.Println("Capacity", cap(slice1))
	fmt.Println("Length", len(slice1))
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		slice1[i] = i
	}

	fmt.Println(slice1)
}
