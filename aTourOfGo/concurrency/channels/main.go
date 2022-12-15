package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/3], c)
	go sum(s[len(s)/3:len(s)*2/3], c)
	go sum(s[len(s)*2/3:], c)
	x, y, z := <-c, <-c, <-c // receive from c

	fmt.Printf("x = %d\ny = %d\nz = %d\nsum = %d", x, y, z, x+y+z)
}
