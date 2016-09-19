package main

import "fmt"

// A rune is an alias for an int32
// runes point to a character list in the ASCII chart
// Ex. 35 = '#'

func main() {

	// Initialize variable 'j' with the value of 35
	j := 35

	// Display the variable 'j' converted to a string
	fmt.Println(string(j))

	// Run a 'for' loop to display characters from 30 - 40
	for i := 30; i <= 40; i++ {
		//value    conv to strng      slice string into bytes to show byte conversions for each character
		fmt.Println(i, " -", string(i), " -", []byte(string(i)))
	}

	// Single quotes references the runes literal
	// 'a' holds a single byte representing a literal (0xc3 0xa4)
	foo := 'a'

	// Print the literal of 'a'
	fmt.Println(foo)

	// Display the data type of foo
	fmt.Printf("%T \n", foo)

	// even though the int32 foo variable is being converted to a rune/
	// it still shows int32, because a rune is an alias for an int32
	fmt.Printf("%T \n", rune(foo))

}
