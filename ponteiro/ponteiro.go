package main

import "fmt"

func main() {
	i := 1
	var p *int = nil

	p = &i //&i get i address to p
	*p++   //*p get value

	fmt.Println("p equals to", *p)
	fmt.Println("i equals to", i)
	fmt.Println("p address", p)
	fmt.Println("i address", &i)

	//Go não permite aritmética de ponteiros
	//p++
}
