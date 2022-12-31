package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/suyash/mynice/helpers"
)

var s = "seven"

func overview() {
	var myStr string
	myStr = "Green"
	fmt.Print("--Pointer--")
	log.Println("my string set to ", myStr)
	changUsingPointer(&myStr)

	log.Println("after func call my string set to ", myStr)

	fmt.Println("--Types and Structs--")
	log.Println("s is ", s)
	saySomething("xxx")
	user := User{
		FirstName: "Traver",
		LastName:  "Sawler",
	}
	log.Println(user.FirstName, user.LastName, user.BirthDate)

	fmt.Println("--Receivers: Structs with Function--")
	var myVar myStruct
	myVar = myStruct{
		FirstName: "John",
	}

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar set to ", myVar.printFirstName())
	log.Println("myVar2 set to", myVar2.printFirstName())

	fmt.Println("--Maps and Slices--")

	myMap := make(map[string]string) //immutable,not sorted,

	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])

	myMapUser := make(map[string]User)

	myMapUser["me"] = User{FirstName: "Travor"}
	log.Println(myMapUser)

	var mySlice []string //slice

	mySlice = append(mySlice, "Travor")
	mySlice = append(mySlice, "Mary")

	log.Println("slice is ", mySlice)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	log.Println("slice number", numbers[0:2])

	fmt.Println("--Loops--")

	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	for i, num := range numbers {
		log.Println(i, num)
	}

	for key, val := range myMap {
		log.Println(key, val)
	}

	fmt.Println("--Interface--")

	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}
	PrintInfo(&dog)

	gorrila := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorrila)

	var myVar3 helpers.SomeType
	myVar3.TypeName = "Some Name"
	log.Println(myVar3)

	fmt.Println("--channels--")
	intChan := make(chan int) //can only hold int
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan

	log.Println(num)

	fmt.Println("--Read / Write from JSON--")
	myJson := `[
		{
			"first_name":"Clark",
			"last_name":"Kend",
			"hair_color":"black",
			"has_dog":true
		},{
			"first_name":"Bruce",
			"last_name":"Wayne",
			"hair_color":"black",
			"has_dog":false
		}
	]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshaling json ", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	var myJsonSlice []Person
	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	myJsonSlice = append(myJsonSlice, m1)

	newJson, err := json.MarshalIndent(myJsonSlice, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))

}

type Person struct {
	FirstName string `json:"first_name`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

const numPool = 10

func CalculateValue(intChan chan int) {
	randomeNumber := helpers.RandomNumber(numPool)
	intChan <- randomeNumber //put inside intChan
}

func PrintText(s string) {
	log.Println(s)
}

func changUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}

func saySomething(s string) (string, string) {
	log.Println("s from the saySomething func is ", s) //taking local value--variable shadowning
	return s, "Worlds"
}

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string { //m is called receiver
	return m.FirstName
}

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4

}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2

}
