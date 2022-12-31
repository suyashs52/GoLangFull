package main

import (
	"fmt"
	"runtime"
)

var x_bool bool

var y float64

func main5() {

	fmt.Println(x_bool)

	x_bool = true

	fmt.Println(x_bool)

	a_int := 56
	b_int := 56

	fmt.Println(a_int == b_int) //a equal to b
	fmt.Println(a_int <= b_int)
	fmt.Println(a_int >= b_int) // a greater or equal to
	fmt.Println(a_int != b_int) //a not equal to b

	fmt.Println(a_int > b_int) // a greater than b
	fmt.Println(a_int < b_int) //a less than b

	fmt.Println(42 == 42)
	//1000

	//numeric type
	fmt.Println("--------Numeric Type Example----")

	x := 42
	y = 42.00

	//x = 2.3456 //float can't be of type int

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	//var x_int8 int8 = -129 //can have negative and positive no, so range is -128 to 127
	var x_int32 int32 = 234
	//fmt.Println(x_int8)
	fmt.Println(x_int32)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println("-----String------") // read only slice/sequence of byte (uint8)
	//string is immutable --> once a string is created it is impossible to change the contents of a string

	s := "Hello , 世界"

	fmt.Printf("\n%s\n", s)
	fmt.Printf("\n%q\n", s)
	fmt.Printf("\n%x\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x %d \n", s[i], s[i])
	}

	fmt.Print("\n--Uni Code char")

	for i := 0; i < len(s); i++ {

	}

	for i, v := range s {
		fmt.Printf("%#U \t %d\n", v, i)
	}

}
