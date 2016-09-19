package main

import "fmt"

func main() {

	for n := range Calcu(Loader(100)) {
		fmt.Println(n)
	}
}

func Loader(in int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++{
			out <- i
		}
		close(out)
	}()
	return out
}

func Calcu(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c{
			out <- n * n
		}
		close(out)
	}()
	return out
}
