package main

import "fmt"

func main() {

	// Assign variable 'a' to 42
	a := 52

	// Display the value of 'a'
	fmt.Println(a)

	// Display the memory address of 'a' by using '&'
	fmt.Println(&a)

	// Assign the variable 'b' to reference the memory address of 'a' by using the '&' sign
	// The *int tells Go that the variable is referencing a memory address where an int is stored
	var b *int = &a // 0xc08203c118

	// Display b which references the memory addres of 'a'
	fmt.Println(b)  // 0xc08203c118
	fmt.Println(*b) //43 - '*b' will dereference

	// Change the value of the variable 'b' is referencing to '42' - in this case variable 'a' will
	// become 42
	*b = 42

	// Display the new result of 'a'
	fmt.Println(a)
}
