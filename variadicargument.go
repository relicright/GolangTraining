package main

import "fmt"

func main() {

	// Initialized a slice variable named 'data' and assigned slices
	data := []float64{43, 56, 87, 12, 45, 57}

	// Assign the returned value of 'average()' to the variable n
	// using '...' at the end of the passed in variable, all items within the variable are passed in
	// in this case - 43, 56, 87, 12, 45, 57 - will all be passed in
	n := averageA(data...)

	// Display 'n' to the console
	fmt.Println(n)
}

func averageA(sf ...float64) float64 {

	// Idiomatic way to declare 'total'
	var total float64

	// If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
	for _, v := range sf {

		// Add each 'range'ed item to total
		total += v
	}

	// return the value of total / the length of sf - in this case 6
	return total / float64(len(sf))
}
