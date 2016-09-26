package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("my-cookie")
		fmt.Println(cookie, err)

		http.SetCookie(w, &http.Cookie{
			Name: "my-cookie",
			Value: "a new value",
		})
	})

	http.ListenAndServe(":8080", nil)
}
