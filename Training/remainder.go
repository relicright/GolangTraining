package main

import "fmt"

func main() {

	// Find the remainder of a division - in this case 3 will go into 12 four times.
	// 3 * 4 = 12.  The remainder would be zero because 3 was able to go into 12
	// enough times that there is no remaining number left over
	x := 12 % 3

	fmt.Println(x)

	// Check if the remainer is an odd or an even number
	if x == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}

	// This works like the remainder above, but 3 is only able to go into 13 four times.
	// This leaves a remainder of '1' since 3 does not go into 13 evenly
	// 3 * 14 = 12 - Since we asked for the remainder of 13 it would then be 13 - 12 = 1
	x = 13 % 3

	fmt.Println(x)

	if x == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}

	// This is a loop that goes through the remainder of 1-20 using '2' as the divisible
	for i := 0; i < 20; i++ {

		// The remainder of  i / 2 would mean the item would be an odd number
		if i%2 == 1 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}

}
