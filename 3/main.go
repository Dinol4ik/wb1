package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func numbersUseWg(numbers [5]int64) {
	wg := &sync.WaitGroup{}
	var sum int64
	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go func(n int64) {
			atomic.AddInt64(&sum, n*n)
			defer wg.Done()
		}(numbers[i])
	}
	wg.Wait()
	fmt.Println(sum)
}
func numbersUseChannel(numbers [5]int64) {
	var sum int64
	c := make(chan int64)
	for i := 0; i < len(numbers); i++ {
		go func(n int64) {
			c <- n * n
		}(numbers[i])
	}
	for i := 0; i < len(numbers); i++ {
		sum += <-c
	}
	fmt.Println(sum)
}
func main() {
	numbers := [5]int64{2, 4, 6, 8, 10}
	numbersUseChannel(numbers)
	numbersUseWg(numbers)
}
