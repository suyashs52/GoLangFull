package main

import "fmt"

// DECLARE the variable x
// ASSIGN the vlaue 21
// declare and assign = initialization
var x = 21

// DECLARE there is a VARIABLE with the IDENTIFIER z
// and that the VARIABLE with the IDENTIFIER z is of TYPE int
// ASSIGN the ZERO VALUE of TYPE int to z
// false for booleans, 0 for intergers, 0.0 for floats, "" for strings,
// and nil for pointers,functions, interfaces,slices, channels and maps
var z int

func main2() {
	//declare a variable and assign a value at the same time
	foovar()
	x = 42
	fmt.Println(x)
	fmt.Println(z)

	x = 99

	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)

	z := "James Bond"
	fmt.Println(z)

}

func foovar() {
	fmt.Println("foovar--", x)
}
