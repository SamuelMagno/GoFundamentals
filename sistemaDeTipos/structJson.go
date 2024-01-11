package main

import (
	"encoding/json"
	"fmt"
)

type itens struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := itens{1, "PC", 5900, []string{"Promocao, Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//json para struct
	var p2 itens
	jsonString := `{"id":2, "nome":"Caneca", "preco":23, "tags":["Casa", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)
}