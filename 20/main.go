package main

import (
	"fmt"
	"strings"
)

func strRevers(str []string) string {
	high := len(str) - 1
	middle := len(str) / 2
	for i := range str {
		if i == middle {
			reversStr := strings.Join(str, " ")
			return reversStr
		}
		str[i], str[high-i] = str[high-i], str[i]
	}
	return ""
}
func main() {
	s := "snow dog sun"
	str := strings.Split(s, " ")
	fmt.Printf("string - %s\n", s)
	res := strRevers(str)
	fmt.Printf("reverse string - %s", res)
}
