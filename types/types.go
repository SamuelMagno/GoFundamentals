package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//integer numerics
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer literal is", reflect.TypeOf(3000))

	//unsigned (only positives) uint, uint16, uint32, uint64

	var b byte = 255
	fmt.Println("A byte is", reflect.TypeOf(b))

	//signed integer int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("Integer maximum value is", i1)
	fmt.Println("'i1' type is", reflect.TypeOf(i1))

	var i2 rune = 'a' //Map the table unicode (int32)
	fmt.Println("'a' unicode is", i2)
	fmt.Println("'i2' type is", reflect.TypeOf(i2))

	//float numerics (float32, float64)
	var x float32 = 49.99
	fmt.Println("x is equal to", x)
	fmt.Println("'x' type is", reflect.TypeOf(x))
	fmt.Println("literal type of 49.99 is", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("'bo' type is", reflect.TypeOf(bo))
	fmt.Println("'bo' is", bo)
	fmt.Println("'!bo' is", !bo)

	//string
	s1 := "Olá, meu nome é Samuel"
	fmt.Println(s1 + "!")
	fmt.Println("String length is", len(s1))

	//type CHAR isn't supported in GO
}
