package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go Foo()
	go Bar()
	wg.Wait()
}

func Foo()  {

	for i:= 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

func Bar()  {
	for i:= 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(45 * time.Millisecond))
	}
	wg.Done()
}
