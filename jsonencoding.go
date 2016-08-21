package main

import (
	//"fmt"
	"encoding/json"
	"os"
)

type Person struct {

	First string
	Last string
	Age int
	notExported int
}

func main()  {

	p1 := Person{"James", "Bond", 20, 007}

	// NewEncoder returns a new *encoder that writes to w.
	//
	//Now that you have an *encoder you can use its receiver function 'Encode()'
	// Encode writes the JSON encoding of v to the stream,
	// followed by a newline character.
	json.NewEncoder(os.Stdout).Encode(p1)
}
