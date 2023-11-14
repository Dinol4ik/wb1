package main

import "fmt"

// грокаем алгоритмы, книга "Aditya Bhargava"

func qSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := make([]int, 0, 1+len(arr)/2)
	greater := make([]int, 0, 1+len(arr)/2)

	for _, val := range arr[1:] {
		if val <= pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	arr = append([]int{}, qSort(less)...)
	arr = append(arr, pivot)
	arr = append(arr, qSort(greater)...)
	return arr
}
func main() {
	fmt.Println(qSort([]int{6, 4, 7, 12, 6, 34, 8, 90}))
}
