package main

import (
	"fmt"
)

// A constant is a simple unchanging value
const p string = "death & taxes"

// You can declare multiple constants at once
const (
	pi = 3.14
	language = "Go"
)

// an iota is a very small quantity
const (
	a = iota  //0
	b = iota  //1
	c = iota  //2
)

// If you declare new iotas in a different declaration they will reset back to 0
const (
	d = iota //0 - Also you do not have to type iota after the first declaration
	e	 //1
	f	 //2
)

// Additional logic can be added to manipulate iotas such a '*'
const (
	g = iota //0
	h = iota * 10 // 1 * 10
	i = iota * 10 // 2 * 10
)

// iotas are capable of more advanced processes
const (
	_ = iota // 0
	//bitwise operations - shifting the bits
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2 * 10)
	gb = 1 << (iota * 10) // 1 << (2 * 10)
)

func main()  {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println("PI - ", pi)
	fmt.Println("Language - ", language)

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)

	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)

	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)

	fmt.Println("thrown away value - ")
	// displayed in binary
	fmt.Printf("kb - %b \n", kb)
	fmt.Printf("mb - %b \n", mb)
	fmt.Printf("gb - %b \n", gb)

	//displayed in decimal
	fmt.Printf("kb - %d \n", kb)
	fmt.Printf("mb - %d \n", mb)
	fmt.Printf("gb - %d \n", gb)

}