package main

import "fmt"

func main() {

	// A for loop can start with an initializing statement
	//init/stmt - condition - Poststmt
	for a := 0; a < 5; a++ {
		fmt.Println(a)
	}

	// A for loop can also evaluate a single condition with no post statement
	b := 6

	// condition
	for b >= 1 {
		fmt.Println(b)
		b--
	}

	// You can nest loops inside of other loops
	for c := 0; c < 5; c++ {
		// Nested loop
		for d := 0; d < 5; d++ {
			println(c, " -", d)
		}
	}

	// 'for' statements can run endlessly as well
	i := 0

	for {
		i++
		// check the remainder of i
		if i%2 == 0 {

			// 'continue' starts the loop from the beginning
			continue
		}

		// Display the value of 'i'
		fmt.Println(i)

		// Check if 'i' becomes more than or equal to 50
		if i >= 50 {

			// use the 'break' keyword to exit the loop
			break
		}
	}
}
