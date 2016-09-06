package main

import (
	"fmt"
	"time"
)

func main() {

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
