package main

import "fmt"

func main() {
	set := [5]string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]int)
	res := make([]string, 0, len(set))
	for _, v := range set {
		if _, ok := m[v]; !ok {
			m[v] = 1
			res = append(res, v)
		}
	}

	fmt.Printf("собственные множества: %v", res)
}
