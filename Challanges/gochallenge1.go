package main

import (
	"fmt"
	"sync"
)

func main() {
	c := gen()

	fw1 := FactWorker(c)
	fw2 := FactWorker(c)
	fw3 := FactWorker(c)
	fw4 := FactWorker(c)
	fw5 := FactWorker(c)
	fw6 := FactWorker(c)
	fw7 := FactWorker(c)
	fw8 := FactWorker(c)

	for n := range merge(fw1, fw2, fw3, fw4, fw5, fw6, fw7, fw8) {
		fmt.Println(n)
	}
}

func gen() <-chan int{
	out := make(chan int)
	go func() {
		for i := 0; i < 1000; i++{
			for j := 3; j < 13; j++{
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func FactWorker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in{
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(c ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	wg.Add(len(c))
	for _, v := range c{
		go func(ch <-chan int) {
			for n := range ch{
				out <- n
			}
			wg.Done()
		}(v)
	}

	go func(){
		wg.Wait()
		close(out)
	}()

	return out
}

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
*/
