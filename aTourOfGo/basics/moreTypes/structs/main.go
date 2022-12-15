package main

import "fmt"

type Zuong struct {
	X int
	Y int
}

type VTD struct {
	name string
	age  int
}

func main() {
	fmt.Println(Zuong{1, 2})
	fmt.Println(VTD{"Dun", 20})
}
