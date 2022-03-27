package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randonNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	//if {init block}; {condition}
	if i := randonNumber(); i > 5 { //switch supported
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Perdeu")
	}
}
