package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		key := "q"

		file, hdr, err := r.FormFile(key)
		fmt.Println(file, hdr, err)

		// We need to set the headers so that the 'Content-Type' displays the page as
		// 'text/html' instead of text/plain.  text/plain will only display your input
		// as text instead of functioning HTML.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, `<form method="POST" enctype="multipart/form-data">

		<input type="file" name="q">
		<input type="submit">

		</form>`)
	})

	//Handle the 'favicon.ico' request to prevent duplicate returns
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
