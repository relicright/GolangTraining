package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup


// init() is a special function that lets you setup initialization functions that
// runs before any code in your scripts
func init()  {
	// You can define the amount of threads to use
	// by using the 'runtime.GOMAXPROCS()' function
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go Foo1()
	go Bar1()
	wg.Wait()
}

func Foo1()  {

	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

func Bar1()  {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
