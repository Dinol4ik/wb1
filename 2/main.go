package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	c := make(chan int)
	result := channel(numbers, c)
	// вывод квадратов с помощью канала
	fmt.Println("Channel--------------------")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-result)
	}
	// вывод квадратов с помощью waitGroup
	fmt.Println("WG--------------------")
	waitGroup(numbers)
}
func waitGroup(numbers [5]int) {
	wg := sync.WaitGroup{}
	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n * n)
			defer wg.Done()
		}(numbers[i])
	}
	wg.Wait()
}
func channel(numbers [5]int, c chan int) chan int {
	for i := 0; i < len(numbers); i++ {
		go func(num int) {
			c <- num * num
		}(numbers[i])
	}

	return c
}
