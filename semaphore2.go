package main

import (
	"fmt"
)

func main() {

	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	//Instead of putting individual channel block state ments such
	// as '<-' , you can loop through the amount of values you want
	// using a 'for' statement.  Here we wait for the number of 'n'
	// loops.  Each loop will take a value off the channel.  Once
	// the number of loops has been reached, the channel will be closed
	// so that the 'range' below can loop over it draining the values
	go func() {
		for i := 0; i < n; i++{
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

