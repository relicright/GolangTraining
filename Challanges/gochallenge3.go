package main

import (
	"fmt"
	"sync"
)

var count int64
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	c1 := incrementor("1")
	c2 := incrementor("2")
	c3 := incrementor("3")

	for m := range mer(c1,c2,c3){
		fmt.Println(m)
	}
	fmt.Println("Final Counter:", count)
}

func incrementor(s string) chan int{
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- i
			fmt.Println("Process: "+s+" printing:", i)
			mutex.Lock()
			count++
			mutex.Unlock()
		}
		close(out)
	}()
	return out
}

func mer(c ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(c))

	for _, v := range c{
		go func(ch chan int) {
			for n := range ch{
				out <- n
			}
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
