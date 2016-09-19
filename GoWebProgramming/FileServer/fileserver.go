package main

import (
	"fmt"
	"text/template"
	"log"
	"net/http"
)

type Page struct {
	Image string
}

func main() {

	tpl, err := template.ParseFiles("templates/index1.html")
	if err != nil{
		log.Fatalln(err)
	}

	//We can use the call to '/resources/' here, but Go will know that the files
	//are located elsewhere by using the 'Handle' below to control the directory
	//flow
	img := Page{`<img style="width:300px" src="/resources/dragonbg.jpg">`}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		if err := tpl.ExecuteTemplate(w, "index1.html", img); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	//We can define a directory to hold our resources and control the way they are received.
	//Here we  use 'Handle()' to handle any requests to the "/resources/" directory.  We then
	//strip the prefix '/resources' away and use a second Handler function to control where
	//the files are being pulled from.  In this case we are pulling from out 'pics' directory
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("pics"))))

	fmt.Println(http.ListenAndServe(":8080", nil))
}




