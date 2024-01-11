package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n) //passando valor de n, não incrementa N original
	fmt.Println(n)

	inc2(&n) //passando referência de n, incrementa N original
	fmt.Println(n)
}
