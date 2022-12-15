package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/* === Sử dụng Waitgroup để biết các goroutine đã hoàn tất === */
/*
 * Để sử dụng waitgroup chúng ta cần import package "sync".
 * Nguyên tắc hoạt động của WaitGroup khá đơn giản là
 * có bao nhiêu goroutines cần đợi thì dùng hàm Add(number).
 * Mỗi khi goroutines chạy xong thì gọi hàm Done().
 * Hàm Wait() sẽ bị block cho tới khi đã done hết tất cả
 * số lượng goroutines đã khai báo trước đó.
 */
func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	wc := new(sync.WaitGroup)
	wc.Add(2)

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		fmt.Println("Goroutine 1 done.")
		wc.Done()
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		fmt.Println("Goroutine 2 done.")
		wc.Done()
	}()

	wc.Wait()
	fmt.Println("All Goroutines done")
}
