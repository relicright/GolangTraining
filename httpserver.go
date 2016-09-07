package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello from the go web server!</h1>")
}

func main() {

	var h Hello

	// ListenAndServe listens on the TCP network address addr
	// and then calls Serve with handler to handle requests
	// on incoming connections.
	// Accepted connections are configured to enable TCP keep-alives.
	// Handler is typically nil, in which case the DefaultServeMux is
	// used.
	err := http.ListenAndServe("localhost:4000", h)
	checkError2(err)


}

func checkError2(e error)  {
	if e != nil {
		log.Fatalln(e)
	}
}
