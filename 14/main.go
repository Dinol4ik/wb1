package main

import "fmt"

func runtimeDetect(t interface{}) {
	switch t.(type) {
	default:
		fmt.Printf("unexpected type %v\n", t)
	case string:
		fmt.Printf("%v - is string\n", t)
	case int:
		fmt.Printf("%v - is int\n", t)
	case bool:
		fmt.Printf("%v - is bool\n", t)
	case float64:
		fmt.Printf("%v - is float64\n", t)
	}
}
func main() {
	i := 1
	runtimeDetect(i)
	f := 2.21
	runtimeDetect(f)
	s := "aaaa"
	runtimeDetect(s)
	b := true
	runtimeDetect(b)
	var c float32
	c = 1.2
	runtimeDetect(c)
}
