package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo string
	turbo  bool
}

func (f *ferrari) ligarTurbo() {
	f.turbo = true
}

func main() {
	car1 := ferrari{"f40", false}
	car1.ligarTurbo()

	var car2 esportivo = &ferrari{"f50", false}
	car2.ligarTurbo()
	
	fmt.Println(car1)
	fmt.Println(car2)
}