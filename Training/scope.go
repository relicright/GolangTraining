package main

import (
	"fmt"
)

// A package level variable can be used throughout the entire scope of the page
var packagelevelvariable string = "package level scope"

func main() {

	// Printing out the package level variable to show it can be used throughout all scopes of the file
	fmt.Printf("The package level variable is: %s\n", packagelevelvariable)

	//The inner block variables can only be used within the block they are assigned
	innerblockvariable := "inner block variable"

	//Printing out the innerblockvariable will work because the variable is withing the scope of the block that is calling it
	fmt.Printf("The inner block level variable is: %s\n", innerblockvariable)
}

// Here is a test function that shows you cannot call a variable that is not within it's scope
func scopevariablefunction() {

	// Uncomment the line below to see the error received when attempting to call a variable that is out of scope
	// fmt.Printf("This is an attempt to call a variable that is out of scope: %s", innerblockvariable)
}

// You cannot call a variable before it is declared.  Below is an example of a bad call
func callbeforedecalre() {

	// Uncomment the lines below to see the effect of trying to call a variable before it is declared
	// fmt.Printf("This is attempting to reference a variable before it is declared: %s", declaredvariable)
	// declaredvariable := "this variable is declared after the item calling it"
}
