package main

import "fmt"

func main()  {

	// Anonymous function
	greeting := func() {
		fmt.Println("hello world")
	}

	// Calling the anonymous function
	greeting()

	// Display the data type of 'greeting'
	fmt.Printf("%T\n", greeting)


	greet := makeGreeter()
	fmt.Println(greet())

}

func makeGreeter() func() string  {
	return func() string {
		return "hello world"
	}
}

