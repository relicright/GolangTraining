package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("https://www.gutenberg.org/files/76/76-0.txt")

	if err != nil {
		log.Fatalln(err)
	}

	// NewScanner returns a new Scanner to read from r.
	byteslice := bufio.NewScanner(res.Body)

	// ALWAYS CLOSE THE RESPONSE
	defer res.Body.Close()

	// ScanWords is a split function for a Scanner that returns each
	// space-separated word of text, with surrounding spaces deleted. It will
	// never return an empty string. The definition of space is set by
	// unicode.IsSpace.
	byteslice.Split(bufio.ScanWords)

	// Scan advances the Scanner to the next token, which will then be
	// available through the Bytes or Text method. It returns false when the
	// scan stops, either by reaching the end of the input or an error.
	// After Scan returns false, the Err method will return any error that
	// occurred during scanning, except that if it was io.EOF, Err
	// will return nil.
	for byteslice.Scan() {
		fmt.Println(byteslice)
	}

}
