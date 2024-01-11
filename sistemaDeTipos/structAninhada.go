package main

import "fmt"

type item struct {
	productId int
	qtde      int
	preco     float64
}

type pedido struct {
	userId	int
	itens 	[]item
}

// Método: Função com receiver (Receptor)
func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	var pedido1 pedido
	pedido1 = pedido{
		userId: 1,
		itens: []item{
			item{productId: 1, qtde: 3, preco: 10.50},
			item{productId: 2, qtde: 2, preco: 23.80},
		},
	}
	fmt.Println(pedido1, pedido1.valorTotal())
}
