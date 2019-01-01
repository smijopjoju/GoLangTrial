package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1, 2, 3)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Println(" Len: %d, Cap: %d content %v", len(s), cap(s), s)
}
