package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Smijo"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	numb := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numb)
}
