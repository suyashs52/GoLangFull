package main

import "fmt"

//goLang is STATIC programing language
//a VARIABLE is declared to hold a VALUE of certain TYPE
//not a DYNAMIC programming language

// package scope
var z1 int = 10
var z2 int
var a_str string = `Shaken,
"Tet,'

, not stirred"`

var z_by byte
var z_ru rune = 100

func main3() {
	fmt.Println(z1)
	z1 = 21

	fmt.Println(z1)
	z1 = 30
	foobar_type()

	fmt.Println(a_str)

	fmt.Printf("Type of z1 %T %T\n", z1, z1)
	fmt.Printf("Type of a_str %T \n", a_str)
	var z_bool bool
	fmt.Println(z1, z2, z_by, z_ru, z_bool)

	//false for boolean
	//0 for integers
	//0.0 for floats
	//"" for Strings
	//nil for
	//pointer,function,interfaces,slices,channels,maps
	//Note: use short declaration operator as much as possible
	//use var for zero value and package scope

}

func foobar_type() {
	fmt.Println(z1)
}
