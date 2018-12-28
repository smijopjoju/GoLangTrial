package main

import (
	"fmt"
	"math"
)

func pow(x, n, limit float64) float64 {

	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Println(" Limit exceed: ", v)
	}

	return limit
}

func main() {
	fmt.Println(pow(5, 3, 100))
}
