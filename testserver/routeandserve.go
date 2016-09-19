package main

import (
	"fmt"
	"text/template"
	"net/http"
)

type Page struct{
	Name string
}

func main() {

	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		p := Page{"Relic"}
		if name := r.FormValue("name"); name != ""{
			fmt.Println(name)
			p.Name = name
		}

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fileserver := http.FileServer(http.Dir("public"))

	http.Handle("/css/", fileserver)
	http.Handle("/pics/", fileserver)

	fmt.Println(http.ListenAndServe(":8080", nil))
}
