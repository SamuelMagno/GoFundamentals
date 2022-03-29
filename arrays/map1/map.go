package main

import "fmt"

func main() {
	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[12345678945] = "Pedro"
	aprovados[12345678923] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 12345678923)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}
