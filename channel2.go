package main

import (
	"fmt"
)

func main() {

	// Channels are made by using the 'make()' function.
	// 'make()' is used to initialize maps, slices and channels
	c := make(chan int)

	// Here we start an anonymous function that loops through
	// i until it reaches 10
	go func() {
		for i := 0; i < 10; i++ {
			// each iteration will pass its value to the channel
			// using the 'c <-' notation.  In this case, 'c' is our channel
			// Once a value is passed to the channel the channel is
			// blocked until the value is taken out of the channel
			c <- i
		}

		// When you close a channel you can no long put values onto the channel.
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}