package main

import "fmt"

// You can define a struct by using the 'type' keyword followed by
// the identifier.  Then the 'struct' keyword.  A struct is an encapsulation
// or aggregate of variables/methods
type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	// You can pass a struct into another struct to achieve
	// a version of inheritance
	Person
	First         string
	LicenseToKill bool
}

func main() {

	// A new variable of type 'DoubleZero' will have access to any
	// of the aggregate items stored in it.  Those items can be set
	// Through literals
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	// A structs items can be modified through the dot notation as well
	p1.Person.Age = 30

	fmt.Println(p1)

}
