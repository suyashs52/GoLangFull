package main

import "fmt"

//constant: a simple , unchanging value
//only exist at compile time

const (
	a_const int     = 42
	b_const float64 = 42.98
	c_const string  = "James Bond"
)

const x_const = 42
const y_const = 42.2

type hotdog int

type hotcat float64

const h_cont hotdog = 45

const r_cont hotcat = 79.76

const (
	_ = iota
	a_iota
	b_iota = iota
	c_iota
	d_iota = iota
	e_iota
	f_iota
)

func constant_demo() {
	fmt.Println("constant demo")
	fmt.Println(x_const)
	fmt.Printf("%T \n", x_const)
	fmt.Println(y_const)
	fmt.Printf("%T \n", y_const)

	fmt.Println(h_cont)
	fmt.Printf("%T \n", h_cont)

	fmt.Println("HotDog Type Int --->", hotdog(x_const))

	fmt.Println(r_cont)
	fmt.Printf("Type of r_cont is %T \n", r_cont)

	//in a function param must be of type hotcat
	// he has float number
	// how will he pass the number in ravali function

	fmt.Println(hotcat(45.56))

	fmt.Println(hotdog(45))

	fmt.Println(a_const)
	fmt.Println(b_const)
	fmt.Println(c_const)

	fmt.Println("-------IOTA-------")
	fmt.Println(a_iota)
	fmt.Println(b_iota)
	fmt.Println(c_iota)
	fmt.Printf("%T\n", a_iota)
	fmt.Println(d_iota)

}
