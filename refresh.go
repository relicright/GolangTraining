package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Person struct {
	First   string
	Last    string
	Age     int
	Species string
}

func (p Person) returnSpecies() string {
	return p.Species
}

type Animal struct {
	First   string
	Last    string
	Age     int
	Species string
}

func (p Animal) returnSpecies() string {
	return p.Species
}

type someBeing interface {
	returnSpecies() string
}

func getSpecies(sb someBeing) string {
	return sb.returnSpecies()
}

func GetRequest(w string) *http.Response {
	res, err := http.Get(w)

	if err != nil {
		log.Fatalln(err)
	}

	return res
}

func ResponseReader(res *http.Response) []string {
	newres, _ := ioutil.ReadAll(res.Body)
	resstring := string(newres)
	stringslice := strings.Fields(resstring)

	return stringslice
}

var input string

func main() {

	responsestring := ResponseReader(GetRequest("http://www.avatarsofsosaria.com"))

	m := make(map[string][]string)

	for _, v := range responsestring {
		if strings.Contains(v, "guild") {
			m[v] = append(m[v], v)
			fmt.Println(v)
		}
	}

	fmt.Println(m)

	p := Person{"James", "Bond", 22, "Human"}
	a := Animal{"Magilla", "Gorilla", 76, "Purple Gorilla"}

	fmt.Println(strings.ToUpper(getSpecies(p)))
	fmt.Println(getSpecies(a))

	for input != "exit" {
		fmt.Println("Enter your name")
		fmt.Scan(&input)

		if input == "someone" {
			for i := 0; i < 200000; i++ {
				fmt.Println(i)
			}
		}

		if strings.ContainsRune(input, 'z') {
			fmt.Println("string contains a Z")
		}
		fmt.Println(input)
	}
}
