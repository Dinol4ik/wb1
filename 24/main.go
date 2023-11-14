package main

import (
	"fmt"
	"wb1/24/point"
)

func main() {
	p1 := point.NewPoints(1, 9)
	p2 := point.NewPoints(2, 20)
	distance := point.Distance{}
	res := distance.DistanceBetweenPoints(p1, p2)
	fmt.Printf("Дистанция между точками = %f", res)
}
