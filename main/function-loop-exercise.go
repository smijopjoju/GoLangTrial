package main

import "fmt"

func Sqrt(x float64) float64 {

	z := 1.0
	z -= z*z - x/(2*z)
	return z
}

func main() {

	for i := 1.0; i < 10; i++ {
		fmt.Println("Square root: ", Sqrt(i))
	}

}
