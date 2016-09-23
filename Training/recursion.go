package main

import "fmt"

// Recursion is when a function calls itself
func factorial(x int) int {

	// Here we wait for 'x' to equal '0' so we won't have an endless loop
	if x == 0 {
		// return 1 to end the function
		return 1
	}
	// in the return statement we are taking 'x' and multiplying it by
	// factoral().  Since we are calling the function we are inside of currently
	// it will run the 'factoral()' function again.
	return x * factorial(x-1)
}

func main() {
	// Display each iteration of the factoral function
	fmt.Println(factorial(4))
}
