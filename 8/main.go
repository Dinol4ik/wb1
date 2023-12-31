package main

import (
	"fmt"
)

func main() {
	var number int64
	var bit, index int
	number = 9999999 // 100011
	bit = 1
	index = 43
	if bit > 1 || bit < 0 || index > 64 {
		return
	}

	// формируем битовую маску
	var mask int64 = 1 << (index - 1)
	if bit == 1 {
		fmt.Printf("           mask: %064b\n", mask)
		fmt.Printf("original number: %064b\n", number)
		// xor
		// 0^0=0
		// 0^1=1
		// 1^0=1
		// 1^1=0
		fmt.Printf("  result number: %064b\n", number^mask)
		return
	}

	fmt.Printf("           mask: %064b\n", mask)
	fmt.Printf("original number: %064b\n", number)
	// and not
	// 0&^0=0
	// 0&^1=0
	// 1&^0=1
	// 1&^1=0
	fmt.Printf("  result number: %064b\n", number&^mask)
}
