package main

import (
	"fmt"
	"runtime"
)

/* === Sử dụng 1 channel để lắng nghe dữ liệu từ nhiều nơi === */
func sender(c chan<- int, name string) {
	for i := 1; i <= 100; i++ {
		c <- 1
		fmt.Printf("%s has sent 1 to channel\n", name)
		runtime.Gosched()
	}
}

func main() {
	myChan := make(chan int)

	go sender(myChan, "S1")
	go sender(myChan, "S2")
	go sender(myChan, "S3")

	start := 0

	for {
		start += <-myChan
		fmt.Println(start)

		if start >= 300 {
			break
		}
	}
}
