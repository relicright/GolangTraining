package main

import (
	"fmt"
	"net/http"
)

func main() {

	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		// Fprint uses a writer that is defined, in this case the http.ResponseWriter
		fmt.Fprintf(w, "Hello, Go Web Development")
	})

	// ListenAndServe listens on the TCP network address addr
	// and then calls Serve with handler to handle requests
	// on incoming connections.
	// Accepted connections are configured to enable TCP keep-alives.
	// Handler is typically nil, in which case the DefaultServeMux is
	// used.
	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}

//**************************** SERVERMUX ******************************************************//
//A ServeMux is essentially a HTTP request router (or multiplexor). It compares incoming
//requests against a list of predefined URL paths, and calls the associated handler for the path
//whenever a match is found.
//
//Handlers are responsible for writing response headers and bodies. Almost any object can be a
//handler, so long as it satisfies the http.Handler interface. In lay terms, that simply means it
//must have a ServeHTTP method with the following signature:
//
//ServeHTTP(http.ResponseWriter, *http.Request)
//
//Go's HTTP package ships with a few functions to generate common handlers, such as FileServer,
//NotFoundHandler and RedirectHandler. Let's begin with a simple but contrived example: