package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body string
}

func main() {

	// ParseGlob will parse all of the files in a directory instead of
	// having to parse each individual file in a list of 'template.ParseFiles()'
	tpl, err := template.ParseGlob("templates/*.html")
	if err != nil{
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		if err := tpl.ExecuteTemplate(w, "index1.html", &Page{"Gopher", "34"}); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}




