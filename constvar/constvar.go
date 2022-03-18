package main

import (
	"fmt"
	"math"
	//m "math" - Package reference
)

func main() {
	const PI float64 = 3.1415
	var raio float64 = 3.2

	area := PI * math.Pow(raio, 2)
	//package reference example
	//area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)

}
