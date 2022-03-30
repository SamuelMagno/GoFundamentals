package main

import "fmt"

func change(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1

	return //clean return (Returns the named vars on function returns)
}

func main() {
	x, y := change(2, 3)
	fmt.Println(x, y)
}
