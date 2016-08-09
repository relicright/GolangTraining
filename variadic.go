package main

import "fmt"

func main() {

	// Assign the returned value of 'average()' to the variable n
	n := average(43, 56, 87, 12, 45, 57)

	// Display 'n' to the console
	fmt.Println(n)
}

func average(sf ...float64) float64  {

	// Display what we have in 'sf' to the console
	fmt.Println(sf)

	// Display the data type of 'sf' to the console
	fmt.Printf("%T\n", sf)

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


// If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
/*
	for key, value := range oldMap {
	    newMap[key] = value
	}
 */


//If you only need the first item in the range (the key or index), drop the second:
/*
	for key := range m {
	    if key.expired() {
		delete(m, key)
	    }
	}
 */


// If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:
/*
	sum := 0
	for _, value := range array {
	    sum += value
	}
 */