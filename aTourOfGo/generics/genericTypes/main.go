package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (root *List[int]) addToRoot(val int) List[int] {
	*root = List[int]{root, val}
	return *root
}

func main() {
	root := List[int]{
		nil,
		1,
	}

	for i := 1; i < 10; i++ {
		root.addToRoot(i)
	}
	fmt.Println(root)
}
