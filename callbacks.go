package main

import "fmt"

// A callback calls back to the function in the params
func visit(numbers []int, callback func(int))  {

	// Loop through the slices of 'numbers'
	for _, n := range numbers {

		// use the callback function in the arguments
		callback(n)
	}
}

func main() {

	// Run the visit function and pass in a func(int) (a function that returns int)
	// The function in this case uses print line using the variable 'n'
	visit([]int{1, 2, 3, 4}, func(n int){fmt.Println(n)})
}


//callback: passing a fun as an argument