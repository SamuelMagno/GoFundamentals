package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	for i, numero := range numeros {
		fmt.Printf("Index: %d - Number: %d \n", i, numero)
	}

	for _, numero := range numeros {
		fmt.Printf("Number: %d \n", numero)
	}
}
