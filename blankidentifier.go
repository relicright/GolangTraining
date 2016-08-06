package main

import (
	"fmt"
)

func main() {

	// Declaring/Assigning a variable that stores the 'fmt.Println()' return in myPrintln
	// fmt.Println will return (n int, err error)
	myPrintln, err := fmt.Println("This is a string")

	//Since the error was captured in the 'err' variable above it must be used to avoid a pollution error
	if err == nil {
		// This fmt.Printf will display the 'n int' return from the captured 'fmt.Println' above
		// You may assume it would return a string, but you must use 'fmt.Sprint' in order to return a 'string'
		// instead of an 'int'
		fmt.Printf("The variable is: %v \n", myPrintln)
	} else {
		fmt.Printf("There was an error: %s\n", err)
	}

	// If you would like to throw away the 'err value so you wouldn't need to use it in your code
	// it can be done using a '_' instead of a variable name
	myPrintln2, _ := fmt.Println("This is a string with a thrown away error variable")

	// Once again, this will only return the 'n int' from the 'fmt.Println' line, but the error variable
	// does not require any additional code since it was thrown away in the declaration process
	fmt.Println(myPrintln2)

}
