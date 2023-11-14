package main

import (
	"fmt"
	"strings"
)

func isRepeat(str string) bool {
	m := make(map[string]struct{})
	r := strings.Split(str, "")
	for _, v := range r {
		v = strings.ToLower(v)
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
func main() {
	fmt.Println("start")
	str := "abcd"
	fmt.Println(isRepeat(str))
}
