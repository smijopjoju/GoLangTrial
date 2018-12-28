package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	fmt.Println(Vertex{1, 2})
	v := Vertex{2, 3}
	v.X = 5
	v.Y = 10
	fmt.Println(v.X, v.Y)
	fmt.Println(v)

}
