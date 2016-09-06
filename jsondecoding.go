package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Firstt      string `json:"First"`
	Last        string
	Age         int
	notExported int
}

func main() {

	var p1 Person

	// strings.NewReader() takes a string argument and returns a '*Reader'
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)

	// NewDecoder returns a new decoder that reads from 'rdr'.
	// Decode reads the next JSON-encoded value from its
	// input and stores it in the value pointed to by '&p1'.
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.Firstt)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
