package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor1("Foo:")
	go incrementor1("Bar:")
	wg.Wait()
	fmt.Println("Final Counter", counter)
}


func incrementor1(s string)  {

	for i := 0 ; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		// 'mutex.Lock() will lock the sync so that other functions have to wait
		// to access a variable until the current function is done with it.  This
		// prevents a race condition where two functions are overwriting each other's
		// variables.
		// Lock locks this section of code.
		// If the lock is already in use, the calling goroutine
		// blocks until the mutex is available.
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)

		//'mutex.Unlock()' will restart the waitgroup where it left off allowing the
		// next function to access the variable.
		mutex.Unlock()
	}
	wg.Done()
}
