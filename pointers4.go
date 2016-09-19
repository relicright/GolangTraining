package main

import "fmt"

// The '*int' in the func overload means the variable points to a stored int
func zero(z *int) {

	// By using the '*' pointer tou are changing the referenced variables value
	// not making a new value somewhere else in memory
	*z = 0

	// Display the memory address of z which is referencing the passed in variable
	// in this case 'x' from the func main()
	fmt.Println("The memory address of 'z' inside of fun zero() is:", z)
}

func main() {

	// Initialize 'x' to 5
	x := 5

	// Display the value of x before using zero()
	fmt.Println("The value of 'x' before using the zero():", x) // x is 0

	// Display the memory address of x
	fmt.Println("The memory address of 'x' inside of main is:", &x)

	// Call the 'zero()' and pass in the memory reference of 'x' by using the '&' sign
	zero(&x)

	// Display the value of x after using zero()
	fmt.Println("The value of 'x' after using the zero():", x) // x is 0
}
