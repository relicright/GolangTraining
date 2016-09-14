package main

import (
	"fmt"
<<<<<<< HEAD
=======
	"runtime"
>>>>>>> 3c80288489eaefc454b5e42ad9761fcc6132791b
	"sync"
	"time"
)

<<<<<<< HEAD
// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for.  Then each of the goroutines
// runs and calls Done when finished.  At the same time,
// Wait can be used to block until all goroutines have finished.
var wg sync.WaitGroup

func main() {
	// Here we add items to the 'WaitGroup' using the 'Add' function
	// built inside of the 'WaitGroup'.  The wait group will wait
	// until '2' items have been declared as 'Done'.
	wg.Add(2)

	/*
	a function executing concurrently with other goroutines in the same
	address space. It is lightweight, costing little more than the allocation
	of stack space. And the stacks start small, so they are cheap, and grow by
	allocating (and freeing) heap storage as required.
	 */
	go Foo()
	go Bar()

	// The process will wait here until the 'WaitGroup' had finished
	// waiting for 2 items
=======
var wg sync.WaitGroup

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go Foo()
	go Bar()
>>>>>>> 3c80288489eaefc454b5e42ad9761fcc6132791b
	wg.Wait()
}

func Foo()  {
<<<<<<< HEAD
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:",i)
		// Here we simulate a wait time between the numbering to show
		// that the two goroutines that are running are running at
		// the same time.
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	// the 'Done()' tells the defined 'WaitGroup' that this item has finished
	// and lets it remove one item from it's wait stack.
=======

	for i:= 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
>>>>>>> 3c80288489eaefc454b5e42ad9761fcc6132791b
	wg.Done()
}

func Bar()  {
<<<<<<< HEAD
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:",i)
		time.Sleep(time.Duration(20 * time.Millisecond))
=======
	for i:= 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(45 * time.Millisecond))
>>>>>>> 3c80288489eaefc454b5e42ad9761fcc6132791b
	}
	wg.Done()
}
