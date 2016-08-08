package main

import "fmt"

// A variable to store a string in
var scaninput string

func main() {

	// Display 'Enter text to scan: ' on the console
	fmt.Println("Enter text to scan: ")

	// Listen for input to scan and put it in the memory address box of 'scaninput'
	// the '&' will put whatever is read in the scan into the memory reference of 'scaninput'
	fmt.Scan(&scaninput)

	// Display the input that was scanned
	fmt.Println(scaninput)
}
