package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string

	// You can prevent strcut variables from being passed the JSON by
	// using the "-" value in the 'json' key
	// Struct values encode as JSON objects. Each exported struct field
	// becomes a member of the object unless
	//   - the field's tag is "-", or
	//   - the field is empty and its tag specifies the "omitempty" option.
	Last string `json:"-"`

	// You can change the identifier passed into the JSON by putting the
	// new identifier as the value in the 'json' key
	Age int `json:"wisom score"`
}

func main() {

	// Create a person
	p1 := Person{"James", "Bond", 20}

	// Catch the slice of bytes in 'bs' created by the 'json.Marshal()'
	// Marshal traverses the value p1 recursively.
	// If an encountered value implements the Marshaler interface
	// and is not a nil pointer, Marshal calls its MarshalJSON method
	// to produce JSON. If no MarshalJSON method is present but the
	// value implements encoding.TextMarshaler instead, Marshal calls
	// its MarshalText method.
	bs, _ := json.Marshal(p1)

	// Print the slice of bytes converted to a string
	fmt.Println(string(bs))
}
