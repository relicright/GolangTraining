package main

import "fmt"

// Create a struct
type Person struct {
	first string
	last  string
	age   int
}

// A receiver allows any struct that uses type 'Person' to call this
// function as if it was embedded
func (p Person) FullName() string {
	return p.first + p.last
}

func main() {

	// Create a 'Person' using identifier 'p'
	p := Person{"Relic", "Right", 31}

	// Print the full name of the 'Person' by calling the receiver function
	fmt.Println(p.FullName())
}
