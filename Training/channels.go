package main

import (
	"fmt"
	"time"
)

func main() {

<<<<<<< HEAD
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
	}()

	go func() {
		for {
			// Here we take the value out of the channel using the
			// '<-' notation which means we are extracting the value
			// from 'c'
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}

=======
	 c := make(chan int)

	go func() {
		for i := 0; i < 20; i++{
			c <- i
		}
		// if you are using range over a channel you need to close the
		// channel before you can range over it
		close(c)
	}()

	go func() {
		for loop := range c {
			fmt.Println(loop)
		}
	}()

	time.Sleep(3 * time.Millisecond)
}
>>>>>>> 3c80288489eaefc454b5e42ad9761fcc6132791b
