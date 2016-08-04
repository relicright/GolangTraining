package main

import (
	"fmt"
)

func main() {

	fmt.Println(42)

	// Print out in binary
	fmt.Print("Print 42 in Binary: ")
	fmt.Printf("%b\n", 42)

	// Print out in Hexadecimal
	fmt.Print("Print 42 in Hexadecimal: ")
	fmt.Printf("%x \n", 42)

	// Print out in Hexadecimal with 0x notation
	fmt.Print("Print 42 in Hexadecimal with 0x notation: ")
	fmt.Printf("%#x \n", 42)

	// Print out in Hexadecimal to ALL CAPS with 0x notation
	fmt.Print("Print 42 in Hexadecimal with 0x notation: ")
	fmt.Printf("%#X \n", 42)

	// Print out in all numeral systems with tab spacers
	fmt.Print("Print out in all numeral systems with tab spacers: ")
	fmt.Printf("%d \t %x \t %#x \n", 42, 42, 42)

	// Loop through the first 200 chars of decimal, binary and hexadecimal
	for i := 0; i < 200; i++ {
		fmt.Printf("Decimal: %d \t Binary: %b \t Hexadecimal: %#x \n", i, i, i)
	}

	// Loop through the first 200 characters of UTF-8
	for i := 0; i < 200; i++ {
		fmt.Printf("UTF-8: %q\n", i)
	}

}
