package main

import (
    "fmt"
    "encoding/json"
    "os"
)

type Person struct {
	First   string
	Last    string
	Age     int
	Species string
}

func (p Person) ListSpecies() string {
	return p.Species
}

type Animal struct {
	First   string
	Last    string
	Age     int
	Species string
}

func (a Animal) ListSpecies() string {
	return a.Species
}

type Being interface {
	ListSpecies() string
}

func retSpecies(b Being) string {
	return b.ListSpecies()
}

func main() {

	p := Person{"James", "Bond", 20, "Human"}
	a := Animal{"Rampage", "Jackson", 4, "Dog"}

	fmt.Println(retSpecies(p))
	fmt.Println(retSpecies(a))

    var storeparse interface
    json.NewEncoder(os.Stdout).Encode(&storeparse)

}
