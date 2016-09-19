package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x) // Address in func zero - base hexadecimal
	fmt.Println(&x)        // address in func zero
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main - base hexadecimal notation - 0x
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}
