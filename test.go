package main

import (
	"fmt"
	"strconv"
)

// Global variable
var newValue int = 900

var (
	name string = "tuanh"
	age  int    = 14
)

func main() {

	// Declare
	var k int = 90
	i := 10

	// Cover type
	var converTypeFloat float32 = float32(k)
	var newString string = strconv.Itoa(k)

	fmt.Println("Hello, 世界", i)
	fmt.Println("Hello, 世界", newValue)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", converTypeFloat, converTypeFloat)
	fmt.Printf("%v, %T\n", newString, newString)

	// Data types
	/*
		Primative
		1. Bool
		2. Numberic
			- Integer
				- signed integer 	int8 int16
				- unsigned integer  uint8 uint16
			- Float float32 float64
			- Complex complex64 complex128
		3. Text
			- String
			- Byte -> UTF-8
			- Rune -> UTF-32
	*/

	// Example
	var complexEx complex64 = 1 + 2i
	var complexEx1 complex64 = complex(1, 7)
	fmt.Printf("%v, %T\n", real(complexEx), real(complexEx))
	fmt.Printf("%v, %T\n", imag(complexEx), imag(complexEx))
	fmt.Printf("%v, %T\n", imag(complexEx1), imag(complexEx1))

	// String
	var name string = "tuanh"
	var arr []byte = []byte(name)
	fmt.Printf("%v, %T\n", arr, arr)
}
