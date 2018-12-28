package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f: ", r)
		}

		fmt.Println(" Inside deferred method of f")
	}()

	fmt.Println("Calling function G")
	g(0)
	fmt.Println(" Returning normally from G")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking !")
		panic(fmt.Sprintf("Initating panicking: ", i))
	}
	defer fmt.Println("Deferred in G: ", i)
	fmt.Println("Normal flow of G")
	g(i + 1) // recursive call
}

func main() {
	f()
	fmt.Println(" Retured normaly from f.")
}
