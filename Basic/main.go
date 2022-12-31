package main

import "fmt"

func main1() {
	fmt.Println("hello team")
	fmt.Println("Hello, every one")
	fmt.Println("Everyone , this is the most awesome class in the entire world for having fun")
	fmt.Println("Having fun and learning go programming languate ")
	fmt.Print("\n\n\n-----SWEET-------\n")
	//control flow
	//sequential
	//iterative
	//conditional
	foo()
	bar()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("VL", i)
		}

	}
	isPrime(10)
	exitMethod()
}

func foo() {
	fmt.Println("I am foo...")
}

func bar() {
	fmt.Println("I am bar")
}

func exitMethod() {
	fmt.Println("I am exiting from the code......:(")
}

func isPrime(n int) bool {
	fmt.Println("ISPRIME", n*10)
	return true
}

///////////////////////////////////
//
// prime no -- self and 1 divisibility , no negative no is prime number , 0 and 1 is not a prime no .2 .......
// // write a function pass a value in that func to check if it prime or not
// to test it with all conditions -- code coverage must be >90%
// write the logic for prime no, and test the prime mehtod with 90%+ coverage
//////////////////////////////////
