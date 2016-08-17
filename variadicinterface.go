package main

import "fmt"

func main() {

	// This is a variable of type 'interface' which mean any type of input can
	// be stored inside of it's slice '[]'
	var varint []interface{}

	// Append some new values to the slice since the slice was never initialized
	// When appending to a slice a new underlying array is created
	varint = append(varint, 42, 3.145, "foo")

	// Run the 'variadicinterface()' function.  When you want to use all the data a slice
	// has you must use the '...' expression after the slice you want to use.  In this case 'variant'
	variadicinterface(varint...)
}

// To make a function that can take in multiple arguments in a single variable you must use the '...' expression
// before the data type you want to use
func variadicinterface(varint ...interface{}) {

	// Range returns (key/index - value)
	for _, v := range varint {

		// Here we check the data type of the index being ranged over
		switch v.(type) {
		case int:
			fmt.Println("this is an int")
		case float64:
			fmt.Println("this is a float64")
		case string:
			fmt.Println("This is a string")
		default:
			fmt.Println("This is the default")
		}
	}
}
