package main

import (
	"fmt"
	"strings"
)

func strRevers(s string) string {
	str := strings.Split(s, "")
	high := len(str) - 1
	middle := len(str) / 2
	for i := range str {
		if i == middle {
			reversStr := strings.Join(str, "")
			return reversStr
		}
		str[i], str[high-i] = str[high-i], str[i]
	}
	return ""
}
func main() {
	s := "▁ ▂ ▃ ▄ ▅ ▆ ▇ █"
	fmt.Printf("string - %s\n", s)
	res := strRevers(s)
	fmt.Printf("reverse string - %s", res)
}
