package main

import "fmt"

func main() {

	// In order to cast an interface type to another
	// type you must use assertion and not conversion
	var val interface{} = "Sydney"
	// var val interface{} = 77

	// You can test a value through assertion.
	// Here if the val is not a string then
	// ok will be false
	newval, ok := val.(string)

	if ok {
		// This shows that 'ok' is of type bool
		fmt.Printf("%T\n", ok)
		fmt.Println(newval)
	}else{
		fmt.Println("Value is not a string")
	}


	var val2 interface{} = 7

	// you can use the '.()' notation to assert a value of the
	// type you need the value to be in order to complete an operation
	fmt.Println(val2.(int) + 6)
}
