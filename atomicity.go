package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor2("Foo:")
	go incrementor2("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor2(s string)  {

	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		// race:
		// counter++

		// no race:
		// atomic provides low-level atomic memory primitives useful for implementing
		// synchronization algorithms
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
