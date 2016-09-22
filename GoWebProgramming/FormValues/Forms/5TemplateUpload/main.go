package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"path/filepath"
	"io"
	"html/template"
)

func main() {

	templates := template.Must(template.ParseFiles("templates/upload.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST"{

			// FormFile returns the first file for the provided form key.
			// FormFile calls ParseMultipartForm and ParseForm if necessary.
			// FormFile Returns - multipart.File, *multipart.FileHeader, Error
			src, _, err := r.FormFile("q")

			//Check Errors
			if err != nil{
				log.Println("there wan an error:", err)
			}
			// Close the file at the end
			defer src.Close()

			// Create creates the named file with mode 0666 (before umask), truncating
			// it if it already exists. If successful, methods on the returned
			// File can be used for I/O;
			// Return a - *File
			//'os.Create()'s file descriptor has a mode of 'O_RDWR'
			//O_RDWR   int = syscall.O_RDWR   // open the file read-write.

			//We are creating a file at the current directory './'
			// then we name the file 'file.txt'

			dst, err := os.Create(filepath.Join("./", "file.txt"))
			//Check Errors
			if err != nil{
				http.Error(w, err.Error(), 500)
				return
			}
			//close file at the end
			defer dst.Close()

			io.Copy(dst, src)
		}

		if err := templates.ExecuteTemplate(w, "upload.html", nil); err != nil{
			http.Error(w, err.Error(), 500)\
			log.Println(err)
			return
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
