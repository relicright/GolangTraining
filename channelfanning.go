package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// Here we assign the channel returned from 'fanIn()' to C
	// The 'fanIn()' takes two 'chan string' arguments and returns
	// a 'chan int' channel
	c := fanIn(boring("Joe"), boring("Ann"))

	// Then we loop 10 times, printing the first 10 values of the channel 'c'
	// the '<-c' will block the program from exiting by waiting for a value
	// before continuing onto the next loop
	for i := 0; i < 10; i++{
		fmt.Println(<-c)
	}

	// Final print showing the program has ended
	fmt.Println("You're both boring; I'm leaving.")
}


func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		// Here we loop indefinitely until the program ends because the
		// loop will continue until the program has ended
		for i := 0; ;i++{
			// Each loop we add a string to our channel 'c'
			c <- fmt.Sprintf("%s %d", msg, i)
			// We wait a random amount of time before the next loop starts
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	// Then return the channel
	return c
}

//Fan IN
// This function takes two arguments that ask for 'chan string' types.
// The, the function returns a 'chan string'
func fanIn(input1, input2 <-chan string) <-chan string  {
	// make a channel to hold the values of the 'chan string'
	c := make(chan string)

	// Start go routines to handle adding values to the new channel by
	// pull values off the channels used as arguments for the function
	go func() {
		for{
			// This take a value off of 'input1' and puts it onto
			// the new channel 'c'
			c <- <-input1
		}
	}()

	go func() {
		for{
			// This take a value off of 'input2' and puts it onto
			// the new channel 'c'
			c <- <-input2
		}
	}()

	//return channel 'c'
	return c
}