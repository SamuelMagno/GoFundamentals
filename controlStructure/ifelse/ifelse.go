package main

import "fmt"

func printResult(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func testResult(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 {
		return "B"
	} else if nota >= 5 {
		return "C"
	} else if nota >= 3 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)

	fmt.Println(testResult(8.3))
	fmt.Println(testResult(4.3))
	fmt.Println(testResult(2.3))
	fmt.Println(testResult(9.3))

}
