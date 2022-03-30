package main

func printApproved(approved ...string) {
	for _, name := range approved {
		println(name)
	}
}

func main() {
	aprovados := []string{"Maria, Predro, Rafael, Guilherme"}
	printApproved(aprovados...)
}
