package main

import "fmt"

type Human struct {
	Name string
	Age  int
	mail string
}
type Action struct {
	Name string
	Age  int
	H    Human
	Human
}

func (h *Human) setName(name string) {
	h.Name = name
}
func (h *Human) setAge(age int) {
	h.Age = age
}
func (a *Action) setActName(name string) {
	a.Name = name
}
func main() {
	act := Action{
		Age:  1,
		Name: "alex",
		H: Human{
			Name: "cj",
			Age:  14,
			mail: "test@mail.ru",
		},
		Human: Human{
			Name: "Ns",
			Age:  23,
			mail: "spiritWinTi2023@dita.ru",
		},
	}
	// обращение к полям структуру Human
	fmt.Println(act.H.Name) // cj
	// можно обратится сразу к полям структуры Human если поле уникальное
	fmt.Println(act.mail) // spiritWinTi2023@dita.ru

	// обращения к методам структуры
	act.setActName("john")
	fmt.Println(act.Name)   // john
	fmt.Println(act.H.Name) // cj
	// обращения к методам встроенной структуры
	act.H.setName("Max")
	fmt.Println(act.H.Name) // max
}
