package main

import (
	"fmt"
)

func switchstatement(value int) string {
	switch value {
	case 0:
		return "zero"
	case 1:
		return "one"
	default:
		return "not zero or one"
	}
}

func main() {
	var someval int
	fmt.Println(someval)
	fmt.Println("test")
	fmt.Print("hello")
	fmt.Print("world")
	fmt.Println(switchstatement(5))

}
