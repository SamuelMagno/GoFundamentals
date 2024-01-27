package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	fmt.Println("Executou")
	c <- 4
	c <- 5
	fmt.Println("Não executou")
	c <- 6
	c <- 7
	c <- 8
	c <- 9
}

func main() {
	c := make(chan int, 3)

	go rotina(c)

	time.Sleep(time.Second)

	fmt.Println(<-c)
}