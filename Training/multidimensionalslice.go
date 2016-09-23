package main

import "fmt"

func main() {

	// A multi-dimensional slice is shown with the '[]' followed by another '[]' or slice operand
	// In this case we are making a slice '[]' of a slice of strings or '[]string'
	// So in other words, we are making a slice of slices
	records := make([][]string, 0)

	//person1
	person1 := make([]string, 3, 3)
	person1[0] = "Todd"
	person1[1] = "Mcleod"
	person1[2] = "42"

	records = append(records, person1)

	//person2
	person2 := make([]string, 3, 3)
	person2[0] = "Dan"
	person2[1] = "The Man"
	person2[2] = "32"

	records = append(records, person2)

	fmt.Println(records)

	/////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println(transactions)
}
