package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"strings"
	"math/big"
)

type Tour struct{
	Name, Price string
}

func main() {

	// Assign a string 'url' to the identifier 'url'
	url := "http://services.explorecalifornia.org/json/tours.php"

	// Run our function that uses an 'http.Get()' to receive a http.Response from
	// the web server's JSON file.
	content := contentFromServer(url)

	// Create a variable to contain the slice of 'Tour' structs that will be returned
	// by the 'toursFromJson()' function
	tours := toursFromJson(content)

	// Begin a loop to loop over the slice of Tour '[]Tour' returned in 'toursFromJson()'
	// function
	for _, tour := range tours {
		// We parse the price as a big.Float for accuracy and formatting
		// 'big.ParseFloat' will return 3 values of which we only want the
		// parsed value.
		// big.ParseFlaot arguments take (price string, base, precision, rounding mode)
		// 'big.ParseFloat returns (float, int, error)
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)

		//Then we use a print format to display the output as we want it.
		// in this case, tour name ($price with 2 decimal positions) then a
		// new line feed
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}
}

func contentFromServer(url string) string {

	// Create a identifier and err variable to catch the return of
	// the 'http.Get()' request
	res, err := http.Get(url)

	// Check for errors
	if err != nil {
		log.Fatalln(err)
	}

	// Defer the close of the response so it is closed after the
	// processing is complete
	defer res.Body.Close()

	// Use 'ioutil.ReadAll' to return the body as an array of bytes '[]byte'
	bytes, err := ioutil.ReadAll(res.Body)
	// Check for errors
	if err != nil {
		log.Fatalln(err)
	}

	// Return the array of bytes converting as a string
	return string(bytes)
}

func toursFromJson(content string) []Tour  {

	// Assign a slice of []Tour to the identifier 'tours'
	tours := make([]Tour, 0, 20)

	// Assign a '*Decoder' to 'decoder' by using the 'json.NewDecoder()'
	// json.NewDecoder requires an io.Reader which is created using the
	// 'strings.NewReader()' function, in which we use the 'content' argument
	decoder := json.NewDecoder(strings.NewReader(content))

	// Token returns the next JSON token in the input stream.
	// At the end of the input stream, Token returns nil, io.EOF.
	//
	// Token guarantees that the delimiters [ ] { } it returns are
	// properly nested and matched: if Token encounters an unexpected
	// delimiter in the input, it will return an error.
	_, err := decoder.Token()

	//Check for errors in the 'decoder.Token()' function
	checkError(err)

	// Assign the 'Tour' struct to the 'tour' identifier
	var tour Tour

	// More reports whether there is another element in the
	// current array or object being parsed.
	for decoder.More() {

		// Decode reads the next JSON-encoded value from its
		// input and stores it in the value pointed to by v.
		err := decoder.Decode(&tour)

		//Check for errors in Decode
		checkError(err)

		// Append the 'tour' variable to the 'tours' slice of '[]Tour'
		tours = append(tours, tour)
	}

	// Return 'tours' which contain a '[]Tour'
	return tours
}

func checkError(e error)  {
	if e != nil{
		log.Fatalln(e)
	}
}
