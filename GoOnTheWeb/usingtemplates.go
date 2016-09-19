package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Page struct {
	Name string
}

func main() {

	// Must is a helper that wraps a call to a function returning (*Template, error)
	// and panics if the error is non-nil. It is intended for use in variable initializations
	// such as
	//	var t = template.Must(template.New("name").Parse("html"))
	// *************************************************************************************
	// ParseFiles creates a new Template and parses the template definitions from
	// the named files. The returned template's name will have the (base) name and
	// (parsed) contents of the first file. There must be at least one file.
	// If an error occurs, parsing stops and the returned *Template is nil.
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		p := Page{Name: "Gopher"}
		if name := r.FormValue("name"); name != ""{
			p.Name = name
		}

		// ExecuteTemplate applies the template associated with t that has the given
		// name to the specified data object and writes the output to wr.
		// If an error occurs executing the template or writing its output,
		// execution stops, but partial results may already have been written to
		// the output writer.
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil{

			// If there was a problem we can use 'http.Error(ResponseWriter, string, code)'
			// to alter the user there was a problem
			//********************************************************
			// Error replies to the request with the specified error message and HTTP code.
			// The error message should be plain text.
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}

// ***********TEMPLATES DOCUMENTATION ******************//
//https://www.golang.org/pkg/html/template


// ****************** TEMPLATE PIPELINE ***************//
// https://godoc.org/text/template