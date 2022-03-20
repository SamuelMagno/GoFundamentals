package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	fmt.Println(int(nota))

	// int to string
	fmt.Println("Teste", strconv.Itoa(123))

	//string to int
	num, _ := strconv.Atoi("123") //num recebe o número, _ é uma forma de ignorar o retorno de erro do segundo parâmetro
	fmt.Println(num - 122)

	//string to bool
	b, _ := strconv.ParseBool("true") // "1" is also true, "0" is false
	fmt.Println("bool", b)
}
