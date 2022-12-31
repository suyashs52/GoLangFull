package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This was the home page")
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := AddValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This was the about page and 2+2=%d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 10.0
	f, err := DivideValues(100.0, 10.0)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func DivideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot devided by zeor")
		return 0, err
	}

	result := x / y
	return result, nil
}

// AddValue add two integers and return the value
func AddValues(x, y int) (int, error) { //can accessbiel from outside package if start with caps letter

	sum := x + y
	return sum, nil
}

// main the the main application
// code>preference> setting> association > add *.tmpl key value is html then type
// in about.tmpl file html:5
func main1() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello World!")

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))

	// })

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
