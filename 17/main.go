package main

import "fmt"

// грокаем алгоритмы, книга "Aditya Bhargava"
func binarySearch(s []int, target int) (int, bool) {
	low := 0
	high := len(s) - 1
	for low <= high {
		middle := (high + low) / 2
		if s[middle] == target {
			return middle, true
		}
		if s[middle] > target {
			high = middle - 1
		}
		if s[middle] < target {
			low = middle + 1
		}
	}
	return 0, false
}
func main() {
	s := []int{1, 2, 5, 10, 34, 44, 69, 100, 134, 145, 200}
	index, ok := binarySearch(s, 69)
	if ok {
		fmt.Printf("number found, index - %d", index)
	} else {
		fmt.Printf("there is no such number")
	}
}
