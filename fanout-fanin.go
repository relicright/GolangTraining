package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	//FAN OUT
	// Distribute the sq work across two goroutines that both read from 'in'
	c1 := sq(in)
	c2 := sq(in)

	//FAN IN
	// Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

// This function generates a channel based on the numbers given
// as arguments
func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// This function does the calculation on the channel arguments.
// the calculations are run in parallel
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in{
			out <- n * n
		}
		close(out)
	}()
	return out
}

// This function combines multiple channels into a single channel.
// This is a 'Fan In' design
func merge(cs ...chan int) chan int {
	out := make(chan int)

	// Create a waitgroup to pause while the channels are combined
	var wg sync.WaitGroup
	fmt.Printf("Type: %T\n", cs)

	// Wait for the amount of items that are in the 'cs' arguments
	wg.Add(len(cs))

	// Start a loop that launches goroutines to simultaneously  to begin
	// loading the new channel with the values the the '[]cs'
	for _, c := range cs{
		// In the go anonymous function we use 'ch chan int' as an argument
		// to enforce the current value we are working on.  When you add
		// arguments in the anonymous functions, you can define its value
		// in the '()' at the end of the anonymous function
		go func(ch chan int) {
			for n := range ch{
				out <- n
			}
			wg.Done()

		// Adding this closure with 'c' allows for the goroutine
		// to take the variable 'c' with it so it does not use the
		// Second value to influence the loop that is already running
		// a goroutine
		}(c)
	}

	// Here we start a goroutine that waits for all of the waitgroups to finish
	// then closes the channel so it can be ranged over
	go func() {
		wg.Wait()
		close(out)
	}()
	// return the new single channel
	return out
}