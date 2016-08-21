package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	// Variables in a struct must have an uppercase letter at the
	// beginning in order for the value to be exported to JSON
	First string
	Last string
	Age int

	// If a variable does not have a capital first letter it will
	// not be exported
	notExported int
}

func main() {

	p1 := Person{"James", "Bond", 20, 007}

	// Marshal traverses the value v recursively.
	// If an encountered value implements the Marshaler interface
	// and is not a nil pointer, Marshal calls its MarshalJSON method
	// to produce JSON.

	//func Marshal(v interface{}) ([]byte, error)
	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
