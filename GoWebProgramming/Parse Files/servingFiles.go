package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Name string
	Age string
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index1.html", "templates/index2.html"))

	//HandleFun functions can be used for different root types to display different templates when the
	// browser requests different directories
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		p := Page{"Random", "41"}
		if name := r.FormValue("name"); name != ""{
			p.Name = name
		}
		if age := r.FormValue("age"); age != ""{
			p.Age = age
		}

		if err := templates.ExecuteTemplate(w, "index1.html", p); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/relic", func(w http.ResponseWriter, r *http.Request){

		p := Page{"Relic", "31"}
		if name := r.FormValue("name"); name != ""{
			p.Name = name
		}
		if age := r.FormValue("age"); age != ""{
			p.Age = age
		}

		if err := templates.ExecuteTemplate(w, "index2.html", p); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
