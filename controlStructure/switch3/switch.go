package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float64, float32:
		return "valor real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(tipo(2.45))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
