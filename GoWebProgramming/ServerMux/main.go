package main

import (
	"net/http"
	"io"
)

type CatHandler int

// A can define a method that will allow our 'CatHandler' to fall under
// the interface of 'Handler' by using the
// 'ServeHTTP(ResponseWrite, *Request)' format

//A Handler responds to an HTTP request.
/*
	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/
func (c CatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	//https://en.wikipedia.org/wiki/List_of_HTTP_header_fields
	// Here we are chaning the headers to tell the page to display
	// as 'text/html' instead of 'text/plain'.  text/html will read
	// the text as if it's in html format, but text/plain will just
	// print out the text in the string.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)

	//https://godoc.org/net/url
	//'*http.Request' contains a variable 'URL' that is of type '*url.URL'
	// we can then use the different variables inside the 'URL' struct
	io.WriteString(w, "<h1>" + r.URL.RawQuery + "</h1>")
}

func main() {

	var cat CatHandler

	//ServeMux is an HTTP request multiplexer. It matches the URL of each
	// incoming request against a list of registered patterns and calls the
	// handler for the pattern that most closely matches the URL.
	mux := http.NewServeMux()

	// Handle registers the handler for the given pattern.
	// If a handler already exists for pattern, Handle panics.
	mux.Handle("/cat/", cat)

	// ListenAndServe listens on the TCP network address addr
	// and then calls Serve with handler to handle requests
	// on incoming connections.
	// In this case it calls our new 'mux' which was passed into
	// 'ListenAndServe()'
	http.ListenAndServe(":8080", mux)
}

