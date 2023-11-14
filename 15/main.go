package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = v[:100] - cрез происходит по кол-ву байт, а не по кол-ву символом
	// проблема может возникнуть когда символ будет
	r := []rune(v)
	justString = string(r[:100]) // тут мы уже срезаем кол-во символов
	fmt.Println(justString, "по символам")
	fmt.Println(v[:100], "по байтам")
}

func createHugeString(i int) string {
	rand.NewSource(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz")
	length := i
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func main() {
	someFunc()
}
