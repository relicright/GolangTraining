package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		key := "q"

		// FormValue returns the first value for the named component of the query.
		// POST and PUT body parameters take precedence over URL query string values.
		// If key is not present, FormValue returns the empty string.

		// In this case, the 'FormValue()' is using the 'key' variable above to return
		// the value of that key
		val := r.FormValue(key)
		fmt.Println("Value: ", val)

		// We need to set the headers so that the 'Content-Type' displays the page as
		// 'text/html' instead of text/plain.  text/plain will only display your input
		// as text instead of functioning HTML.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, `<form method="GET">

		<input type="text" name="q">
		<input type="submit">

		</form>`)
	})

	//Handle the 'favicon.ico' request to prevent duplicate returns
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
