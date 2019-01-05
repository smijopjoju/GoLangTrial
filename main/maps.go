package main

import "fmt"

type Vertex struct {
	x int32
	y float32
}

func main() {
	var m map[string]string = make(map[string]string)

	m["smijo"] = "smijo"

	fmt.Println(m)

	var vertexMap map[string]Vertex = make(map[string]Vertex)

	vertexMap["test"] = Vertex{10, 20.546}

	fmt.Println(vertexMap)

	// map literal trial
	type Vertex1 struct {
		Lat, Long float64
	}

	var m1 = map[string]Vertex1{
		"Test1": {5.2, 10.5},
		"Test2": {20.8, 25.4},
	}

	fmt.Println(m1)

	// Map of map trial
	var innerMap map[string]string = make(map[string]string)
	var mapOfMap map[string]map[string]string = make(map[string]map[string]string)

	innerMap["inner"] = "Map"
	mapOfMap["outter"] = innerMap

	fmt.Println("Map of Map Example", mapOfMap)
}
