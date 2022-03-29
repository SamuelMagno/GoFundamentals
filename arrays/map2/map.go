package main

import "fmt"

func main() {
	funcSalario := map[string]float64{
		"José":    1110.45,
		"Gabriel": 3232.90,
		"Pedro":   2323.89,
	}
	funcSalario["Rafael"] = 1345.0

	fmt.Println(funcSalario)
	delete(funcSalario, "Null")

	for nome, salario := range funcSalario {
		fmt.Println(nome, salario)
	}
}
