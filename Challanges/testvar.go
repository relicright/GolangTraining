package main

import "fmt"

func main() {

	var numbers = []int {22, 24, 26, 28, 30}

	fmt.Println(testFunc(numbers...))
}
func testFunc(x ...int) int {

	var total int

	for _, val := range x {
		total += val
	}

	return total
}
