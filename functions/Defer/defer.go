package main

import "fmt"

func approvedValue(num int) int {
	defer fmt.Println("Fim")
	if num >= 5000 {
		fmt.Println("High value")
		return 5000
	}
	fmt.Println("Low value")
	return num
}

func main() {
	fmt.Println(approvedValue(7000))
	fmt.Println(approvedValue(3500))
}
