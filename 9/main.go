package main

import (
	"fmt"
)

func numbers(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// читает числа из канала, возводит в квадрат и передает в другой канал
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	c := numbers(2, 3, 5, 10, 45, 7, 5, 134, 32, 15, 63, 23, 67, 89, 43)
	out := sq(c)

	// выводим результат из выходного канала
	for v := range out {
		fmt.Println(v)
	}
}
