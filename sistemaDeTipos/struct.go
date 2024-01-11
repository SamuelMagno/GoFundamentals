package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: Função com receiver (Receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    2.00,
		desconto: 0.08,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Caneta", 5.00, 0.10}
	fmt.Println(produto2, produto2.precoComDesconto())
}