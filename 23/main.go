package main

import "fmt"

func deleteElemSlice(s []int, index int) []int {
	if cap(s) < index {
		fmt.Printf("элемента с индексом %d не существует\n", index)
		return s
	}
	s = append(s[:index], s[index+1:]...)
	fmt.Printf("элемент под индексом %d удален\n", index)
	return s
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Оригинальный слайс %d\n", s)
	s = deleteElemSlice(s, 9)
	fmt.Printf("Измененный слайс %d", s)
}
