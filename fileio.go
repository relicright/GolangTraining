package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main() {

	err := ioutil.WriteFile("./newfile.txt", []byte("Hello world"), 0644)
	checkError(err)

	fileName := "./newfile.txt"

	content, err := ioutil.ReadFile(fileName)
	checkError(err)

	fmt.Println(string(content))
}

func checkError(e error)  {
	if e != nil{
		log.Fatalln(e)
	}
}
