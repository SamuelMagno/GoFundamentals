package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(2*time.Second)
	c <- 1 //envia 1 dado
}

func main() {
	c := make(chan int)
	go rotina(c)

	fmt.Println(<-c)//recebe 1 dado
	fmt.Println("Foi lido")
	fmt.Println(<-c) //deadlock
	fmt.Println("Fim")

}