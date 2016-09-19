package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := fanInFunc(loadChannel("Mary"), loadChannel("Bob"))
	for i := 0; i < 10; i++{
		fmt.Println(<-c)
	}
}

func loadChannel(words string) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; ;i++{
			out <- fmt.Sprintf("%s %d", words, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return out
}

func fanInFunc(input1, input2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-input1
		}
	}()

	go func() {
		for {
			out <- <-input2
		}
	}()
	return out
}
