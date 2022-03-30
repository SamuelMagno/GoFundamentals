package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Println(media(2.3, 4.5, 9.8))
	fmt.Println(media(7.0, 8.5, 7.8, 0))
	fmt.Println(media(5.2, 6.6, 4.8, 0, 8.5))
}
