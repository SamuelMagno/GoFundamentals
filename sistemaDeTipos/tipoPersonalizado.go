package main

import "fmt"

type nota float64

func (n nota) entre(inicio float64, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func main() {
	var n nota
	n = 7.4

	fmt.Println(n.entre(7,8))
}