package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* === Sử dụng Select để chọn Channel đã sẵn sàng === */
func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		ch2 <- 2
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("Ch1 come first with value:", v1)
		fmt.Println("then ch2 with value:", <-ch2)
	case v2 := <-ch2:
		fmt.Println("Ch2 come first with value:", v2)
		fmt.Println("then ch1 with value:", <-ch1)
	}
}
