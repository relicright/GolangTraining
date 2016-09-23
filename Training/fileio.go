package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"io/ioutil"
)

func main() {

	content := "Hello from go!"

	// To create a file with Go:
	// identifier, error, then os.Create(path to your file)
	file, err := os.Create("./fromString.txt")
	// Check for any errors while creating the file
	CheckError(err)

	// Defer a 'file.Close()' statement so the stream is closed
	// at the end
	defer file.Close()

	//Write to a file
	// identifier, error.  Then, name of the file, content to enter
	ln, err := io.WriteString(file, content)
	// check for an error writing to the file
	CheckError(err)

	// convert 'content' to a slice of bytes '[]byte'
	bytes := []byte(content)

	// you can write to a file using an array of bytes by using the
	// ioutils package.  WriteFile(Filename, data, permissions)
	ioutil.WriteFile("./fromBytes.txt", bytes, 0644)

	fmt.Println("All done with file of %v characters", ln)

}

func CheckError(e error)  {
	if e != nil{
		log.Fatalln(e)
	}
}
