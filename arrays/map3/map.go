package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 1321.78,
			"Gilson":  5434.90,
		},
		"J": {
			"Jose": 1232.54,
			"Juan": 5090.90,
		},
		"P": {
			"Paula": 6312.90,
			"Pedro": 2321.90,
		},
	}

	delete(funcPorLetra, "J")
	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}
}
