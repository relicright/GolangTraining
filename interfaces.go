package main

import (
	"fmt"
	"math"
)

// Create two structs
type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

// Create receiver functions for both structs
// each receiver must have the same name and return type
// in order to be compatible in the same interface
func (z Square) area() float64 {
	return z.side * z.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Create an interface.  The interface determines which structs are
// compatible with the interface by defining the required functions
// within withs braces - in this case, a struct must have a
// method 'area()' with a return type 'float64'
type Shape interface {

	// This is where the methods required are defined
	area() float64
}

// Now in a new function we set the parameters to take a 'Shape'
// Since 'Square' and 'Circle' both implement the 'area()' function,
// They are both eligible to be used as a 'Shape' interface.
func info(z Shape)  {

	fmt.Println(z)
	fmt.Println(z.area())
	fmt.Printf("%T\n", z)

}

func main() {

	// Assign a variable the 'Square' struct
	s := Square{10}

	// Assign another variable the 'Circle' struct
	c := Circle{10}

	// enter the 'Square' struct as the 'Shape' argument required
	// by the 'info()' function
	info(s)

	// enter the 'Circle' struct as the 'Shape' argument required
	// by the 'info()' function
	info(c)
}
