package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func(){
		for i:=1; i<=3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()
	return c
}

// multiplexar
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select{
				case s:= <-entrada1:
					c <- s
				case s:= <-entrada2:
					c <- s
			}
		}
	}()

	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}