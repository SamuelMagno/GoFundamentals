package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "1234", 3)
	//fale("João", "5678", 1)

	//go fale("Maria", "Ei", 30)
	//go fale("João", "Opa", 30)
	//time.Sleep(time.Second*5)
	//fmt.Println("Fim")

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns", 5)
	//A goroutine tem sua execução interrompida quando o programa principal termina

}