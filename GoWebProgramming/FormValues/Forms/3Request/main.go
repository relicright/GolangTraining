package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//We can check what type of 'Method' if being used in our form
		// to determine what steps should be taken to process the information
		if r.Method == "POST"{

			key := "q"

			// FormFile returns the first file for the provided form key.
			// FormFile calls ParseMultipartForm and ParseForm if necessary.
			//This file will be found in the 'POST' data once a file is uploaded
			//using the form below
			file, _, err := r.FormFile(key)

			if err != nil{
				log.Println("There was an error", err)
			}

			//Make sure to close the file at the end of it's use
			defer file.Close()

			// Copy copies from src to dst until either EOF is reached
			// on src or an error occurs.  It returns the number of bytes
			// copied and the first error encountered while copying, if any.

			//Write the file data to the writer.  In this case:
			//We are writing to the 'http.ResponseWriter'
			io.Copy(w, file)
		}

		// Set the content type so the 'WriteString()' below can use the text as
		//functioning HTML
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		//WriteString to our writer to provide a simple upload form in HTML
		// 'enctype="multipart/form-data"' tells the browser that we will be
		//posting a file from our form
		io.WriteString(w, `<form method="POST" enctype="multipart/form-data">

		<input type="file" name="q">
		<input type="submit">

		</form>`)
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
