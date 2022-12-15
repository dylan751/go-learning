package main

import (
	"fmt"
	"math"
)

func main() {
	defer fmt.Println("world")
	defer fmt.Println("world 2")

	fmt.Println("hello")

	a := 3
	b := 4
	c := math.Sqrt(float64(a*a + b*b))
	fmt.Println(c)
}
