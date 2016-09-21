package main

import (
	"fmt"
	"sort"
)

func main() {

	//SLICES ARE DONE BY ADDING A '[]' WITH NO NUMBER THEN A DATATYPE - IN THIS CASE 'STRING'
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	//YOU CAN APPEND ITEMS TO THE LIST BY USING THE 'APPEND()' FUNCTION, THEN THE SLICE TO APPEND 'COLORS', THEN WHAT TO APPEND 'PURPLE'
	//             slice     appended
	colors = append(colors, "purple")
	fmt.Println(colors)

	// YOU CAN REMOVE SLICES BY USING THE '[]' WITH THE SLICE REFERENCE [START : FINISH]
	//             slice  start    finish
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	//YOU DONT NEED TO ADD A START IF YOU WANT TO START FROM THE BEGINNING - IN THIS CASE [0 : LEN(COLORS) - 1] OR EVERYTHING BUT THE LAST ITEM
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	//YOU CAN 'MAKE' A SLICE WITH A PREDEFINED SIZE LIKE AN ARRAY, BUT IT CAN BE CHANGED AT RUNTIME
	//          datatype-slices-capacity
	numbers := make([]int, 5, 5)
	numbers[0] = 122
	numbers[1] = 23
	numbers[2] = 34
	numbers[3] = 4
	numbers[4] = 52
	fmt.Println(numbers)

	//WHEN APPENDING MORE VARIABLES THAT EXCEED THE CAPACITY OF YOUR SLICE, THE SLICE WILL ADD MORE CAPACITY ON THE FLY
	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	//YOU CAN SORT SLICES IN MANY DIFFERENT WAYS USING THE 'SORT' PACKAGE
	sort.Ints(numbers)
	fmt.Println(numbers)

	// Append a slice to another slice
	// Create your first slice of 'int'
	slice1 := make([]int, 5, 50)
	slice1[0] = 0
	slice1[1] = 1
	slice1[2] = 2
	slice1[3] = 3
	slice1[4] = 4

	// Make a new slice to add to 'slice1'
	newslice := []int{5, 6, 7, 8, 9}

	// You can append one slice to another by adding the '...' epression at the end of the variavle name
	slice1 = append(slice1, newslice...)

	fmt.Println("New slice added to 'slice1':\n", slice1[:])

	//You can append slices at their specific index using the '[:]' indexing
	mySlice := []string{"monday", "tuesday"}
	myOtherSlice := []string{"wednesday", "thursday", "friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)

}
