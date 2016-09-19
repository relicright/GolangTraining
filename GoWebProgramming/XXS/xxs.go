package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Heading string
	Input string
}

func main() {

	home := Page{
		Title: "Nothing Escaped",
		Heading: "Nothing is escaped with text/template",
		Input: `<script>alert("Yow!");</script>`,
	}

	// ParseGlob will parse all of the files in a directory instead of
	// having to parse each individual file in a list of 'template.ParseFiles()'
	tpl, err := template.ParseFiles("templates/index1.html")
	if err != nil{
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		if err := tpl.ExecuteTemplate(w, "index1.html", home); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	http.Handle("/resources", http.StripPrefix("/resources", http.FileServer(http.Dir("pics"))))

	fmt.Println(http.ListenAndServe(":8080", nil))
}




