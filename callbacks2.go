package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {

	var xs []int

	// go over each value in numbers using range
	for _, val := range numbers {

		// Callback to the passed in function which returns a bool
		// (n > 1) or (val > 1)
		if callback(val) {

			// If 'val' is greater than '1'
			// append the value of 'val' to variable 'xs'
			xs = append(xs, val)
		}
	}

	// return the slices of int
	return xs
}

func main() {
	// This is capturing the returned slice of ints from the function 'filter'
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {

		// the callback returns a bool - if n > 1 = true/false
		return n > 1
	})

	// Print the slices of int ([]int) to the console
	fmt.Println(xs)
}
