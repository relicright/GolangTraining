package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - \n", &a)
	fmt.Printf("In decimal: %d\n", &a)

	var meters float64
	fmt.Print("Enter meters swam:")
	scanCatch, err := fmt.Scan(&meters)

	if err == nil {
		fmt.Println(&scanCatch)
		yards := meters * metersToYards
		fmt.Println(meters, " meters is ", yards, " yards.")
	} else {
		fmt.Printf("That is not a valid number.  we are looking for a %T", meters)
	}

}
