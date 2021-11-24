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
	fmt.Printf("%v, %T", k, k)
	fmt.Printf("%v, %T", converTypeFloat, converTypeFloat)
	fmt.Printf("%v, %T", newString, newString)
}
