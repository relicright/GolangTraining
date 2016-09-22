package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Person struct{
	First string
	Last string
}

func main() {

	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var p Person
		if r.Method == "POST"{
			p = Person{
				r.FormValue("first"),
				r.FormValue("last"),
			}
		}else{
			p = Person{"none", "none"}
		}

		//w.Header().Set("Content-Type", "text/http; charset=utf-8")

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
