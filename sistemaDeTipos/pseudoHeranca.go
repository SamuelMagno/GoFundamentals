package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

type ferrari struct {
	carro //campo anonimo
	turbo bool
}

func main() {
	f := ferrari{}
	f.nome = "f40"
	f.turbo = true
	f.velocidade = 0

	fmt.Println(f)
}