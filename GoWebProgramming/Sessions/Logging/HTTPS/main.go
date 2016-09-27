package main

import (
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(`This is an example server.\n`))
}

func main() {

	http.HandleFunc("/", handler)
	log.Println("Go to https://localhost:10443/ or https://127.0.0.1:10433/")

	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil{
		log.Fatal(err)
	}
}

//Go can generate certificates & keys by using the command below:

//********************WINDOWS******************************//
//go run %GOROOT%/src/crypto/tls/generate_cert.go --host=localhost




//********************UNIX**********************************//
//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
