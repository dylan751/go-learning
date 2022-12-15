package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

type Person struct {
	name string
	age  int
}

var m map[string]Vertex
var z map[int]Person

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
