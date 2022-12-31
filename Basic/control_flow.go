package main

import "fmt"

func control_flow() {
	fmt.Println("Control Flow")

	//SEQUENTIAL Control Flow
	//hit a LOOP control flow --ITERATIVE control flow
	//CONDITIONAL Control Flow if switch

	// for init; condition ; post {
	// 	fmt.Println()
	// }

	// for i := 0; i <= 100; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i <= 10; i++ {
	// 	for j := 0; j < 3; j++ {

	// 		fmt.Printf("The Outer loop: %d\t The inner loop %d \n", i, j)

	// 	}
	// 	fmt.Println("----")

	// }

	// for inint; condition; post {}
	// for condition {}
	//for {}

	//x := 1
	// for x < 10 {
	// 	fmt.Println(x)
	// 	x++

	// }

	// for {

	// 	if x > 8 {
	// 		break

	// 	}
	// 	fmt.Println(x)
	// 	x++

	// }

	//x = 83 / 40
	// y := 83 % 40

	// fmt.Println(x, y)

	// for {
	// 	x++
	// 	if x > 100 {
	// 		break
	// 	}
	// 	if x%2 == 0 { //odd number 3%2=1!=0
	// 		continue
	// 	}

	// 	fmt.Println(x)
	// }

	// if true {
	// 	fmt.Println("001")
	// }
	// if false {
	// 	fmt.Println("002")
	// }

	// if !true {
	// 	fmt.Println("003")
	// }

	// if !false {
	// 	fmt.Println("004")
	// }

	// if 2 == 2 {
	// 	fmt.Println("005")
	// }

	// x := 41

	// if x == 40 {
	// 	fmt.Println("our value was 40")
	// } else if x == 41 {
	// 	fmt.Println("our value was   41")
	// } else {
	// 	fmt.Println("our value was not 40 or 41")
	// }

	// switch {
	// case false:
	// 	fmt.Println("this should not print")
	// case (2 == 4):
	// 	fmt.Println("this should not print 2")
	// case (3 == 3):
	// 	fmt.Println("it ll print")
	// 	fallthrough
	// case (4 == 4):
	// 	fmt.Println("also true")
	// 	fallthrough
	// case (7 == 9):
	// 	fmt.Println("not true1")
	// 	fallthrough
	// case (11 == 14):
	// 	fmt.Println("not true 2")
	// 	fallthrough
	// default:
	// 	fmt.Println("this is default")
	// }

	//Array: a numberered sequence of element of a single type
	// size of array doesnt not changed

	var x_arr [5]int
	x_arr[4] = 10
	fmt.Println(x_arr)
	fmt.Println(len(x_arr))

	var x_slice []int = []int{4, 5, 6, 7, 34}

	fmt.Println(x_slice)
	fmt.Println("existed")
}
