package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		http.SetCookie(w, &http.Cookie{
			Name: "my-cookie",
			Value: "This is a cookie",
		})
	})

	http.ListenAndServe(":8080", nil)
}
