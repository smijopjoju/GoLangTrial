package main

import (
	"fmt"
	"math"
)

func compute(x1 float64, y1 float64, paramFun func(X, Y float64) float64) float64 {
	return paramFun(x1, y1)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x + y)
	}

	fmt.Println(hypot(10.0, 5.0))
	fmt.Println(compute(10.0, 3.0, hypot))
	fmt.Println(compute(10.0, 2.0, math.Pow))
}
