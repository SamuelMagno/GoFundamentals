package main

import "fmt"

func main() {
	fmt.Println(testResult(10))
	fmt.Println(testResult(7.8))
	fmt.Println(testResult(11))
}

func testResult(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 3, 4:
		return "D"
	case 1, 2:
		return "E"
	default:
		return "Nota inválida"
	}
}
