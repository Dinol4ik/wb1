package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	// возьмем большое число сиз строки
	a, ok := big.NewInt(0).SetString("7454514214644644854341341544748443215648649", 10)
	if !ok {
		fmt.Println("ошибка в создании первого числа")
		os.Exit(1)
	}
	b, ok := big.NewInt(0).SetString("2137132456768715316477934522251418181918141", 10)
	if !ok {
		fmt.Println("ошибка в создании первого второго числа")
		os.Exit(1)
	}

	fmt.Printf("1 число: %s\n2 число: %s\n\n", a.String(), b.String())

	res := big.NewInt(0)
	// алгебрические операции
	fmt.Println("Умножение:", res.Mul(a, b))
	fmt.Println("Деление:", res.Div(a, b))
	fmt.Println("Сумма:", res.Add(a, b))
	fmt.Println("разность:", res.Sub(a, b))
}
