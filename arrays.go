package main

import (
	"fmt"
)

func main() {

	// You can declare an array by starting with the identifier 'colors' and a '[]' with the number of items
	// in the array '5'.  Then add the data type after 'string'
	var colors [3]string

	// You can assign each item in the array by using the identifier - then the '[]' with the item number '0'
	// followed by the '=' operator and then the assignment 'Red'
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	// Display each item in the array separated by a space
	fmt.Println(colors)

	// Display only one item in the array by using the identifier followed by '[]' with the array index inside '1'
	fmt.Println(colors[1])

	// ARRAYS CAN BE SET USING THE '{}' AT THE END OF THE STATEMENT
	var numbers = [5]int{5, 3, 1, 2, 4}

	// Display each item of the array set above
	fmt.Println(numbers)

	//LEN WILL TELL YOU HOW MANY ITEMS ARE IN THE ARRAY
	fmt.Println("number of colors", len(colors))
	fmt.Println("number of numbers", len(numbers))
}
