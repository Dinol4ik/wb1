package main

import "fmt"

func main() {
	set1 := [6]int{1, 2, 3, 4, 8, 5}
	set2 := [6]int{1, 4, 2, 9, 10, 8}
	result := make([]int, 0, 10)

	setStorage := make(map[int]int)
	for _, value := range set1 {
		setStorage[value] += 1
	}
	for _, val := range set2 {
		if setStorage[val] == 1 {
			result = append(result, val)
		}
	}

	fmt.Println(result)
}
