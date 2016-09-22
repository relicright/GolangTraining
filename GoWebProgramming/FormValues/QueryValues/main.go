package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Define the key we want to pull the value from
		key := "q"

		// Query parses RawQuery and returns the corresponding values.
		// In this case we take the RawQuery and pull the value from it.
		// A Query is located in the URL of the web browser usually as a
		// ?key=value     type of statement
		value := r.URL.Query().Get(key)
		io.WriteString(w, "This is the value:" + value)

		//*****************************************************************//

		// Since 'Query' returns a 'map[string][]string' it can be ranged
		// over to pull each of the element values
		for key, rangedValues := range r.URL.Query(){
			for _, v := range rangedValues{
				io.WriteString(w, key + " " + v + "\n")
			}
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
