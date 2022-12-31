package main

import "fmt"

var a int = 10
var b string = "James Bond"
var c bool
var d bool = true

type anjan int

var h1 anjan

type ravali string

var h2 ravali

func main4() {

	e := 42
	f := "Shaken not stirred"
	g := `Ravel says "Hello Anjan"`
	h := `Q says,
	"I have some new toys"
	`

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

	fmt.Println(d)

	fmt.Println(e)

	fmt.Println(g)

	fmt.Println(b, "says, ", f)

	fmt.Println(h)

	fmt.Printf("%v\t", a)

	fmt.Printf("\n+ sign %+v\n", -10)
	fmt.Printf("%b\n", a)
	fmt.Printf("%x\n", a)
	a = 911
	fmt.Printf("%#x\n", a)
	fmt.Printf("%#v\t", a)

	h1 = 100

	fmt.Println("\n", h1)

	h2 = "hello world"

	fmt.Println(h2)

	fmt.Printf("%T\n", h2)

	k100 := int(h1)

	fmt.Println(k100)
	fmt.Printf("%T", k100)

}
