package main

import "fmt"

func bit_shifting() {

	fmt.Println("Bit Shifting")

	s := 6

	fmt.Printf("%d\t\t%b\n", s, s)

	s = s << 1

	fmt.Printf("%d\t\t%b\n", s, s)

	s = s >> 2

	fmt.Printf("%d\t\t%b\n", s, s)

	fmt.Println("----4 demo----")
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	fmt.Printf("Right Shift %d\t\t%b\n", x<<1, x<<1)
	fmt.Printf("Left  Shift %d\t\t%b\n", x>>1, x>>1)

}

//
//KB 1024*byte
//MB 1024* 1kb
//GB
//TB
