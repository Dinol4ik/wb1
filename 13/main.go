package main

import "fmt"

func main() {
	a, b := 2, 10
	fmt.Printf("befor swap a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("after swap a = %d, b = %d", a, b)
}
