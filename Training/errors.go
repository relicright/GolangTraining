package main

import (
	//"fmt"
	"errors"
	"log"
)

func main() {

	//Catch the error from function 'sqrt()'
	_, err := sqrt(-10)

	//Display the error if one is returned
	if err != nil{
		log.Fatalln(err)
	}
}

// sqrt will hold our intentional error to display
func sqrt(f float64) (float64, error) {
	if f < 0 {
		//errors.New will be out created error with a custom string
		return 0, errors.New("Some error here saying about how no square of less than 0")
	}

	// If the number is not below zero, then we return a random number along with nil
	return 42, nil
}
