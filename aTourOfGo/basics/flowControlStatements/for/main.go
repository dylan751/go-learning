package main

import "fmt"

func main() {
	// -------------- Normal for --------------
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// -------------- for is Go's "while" --------------
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// -------------- infinite loop --------------
	for {
		fmt.Printf("Hello")
	}

}
