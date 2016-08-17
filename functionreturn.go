package main

import "fmt"

func main() {

	fmt.Println(greet("Jane ", "Doe"))

	fmt.Println(greet2("Jane ", "Doe"))

	fmt.Println(greet3("Jane ", "Doe "))
}

//Single return
func greet(fname, lname string) string {

	//
	return fmt.Sprint(fname, lname)
}

// Named return value
func greet2(fname, lname string) (s string) {

	s = fmt.Sprint(fname, lname)
	return
}

// Return multiple values
func greet3(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
