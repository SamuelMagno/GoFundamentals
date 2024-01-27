package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 //Enviando dados
	<-ch    //Recebendo dados

	ch <- 2
	fmt.Println(<-ch)

}