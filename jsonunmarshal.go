package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// This tag shows that json can be unmarshal() by using the json tag
	// as the identifier
	Firstt string `json:"First"`
	Last   string
	Age    int
}

func main() {
	var p1 Person
	//fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	// Simulated JSON marshaled []byte return
	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)

	// Unmarshal parses the JSON-encoded data and stores the result
	// in the value pointed to by '&p1'.
	// To unmarshal JSON into a struct, Unmarshal matches incoming object
	// keys to the keys used by Marshal (either the struct field name or its tag),
	// preferring an exact match but also accepting a case-insensitive match.
	// Unmarshal will only set exported fields of the struct.
	json.Unmarshal(bs, &p1)

	fmt.Println("------------")
	fmt.Println(p1.Firstt)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T\n", p1)
}
