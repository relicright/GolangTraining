package main

import (

)
import (
	"net/http"
	"io"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func getCode(data string) string {

	// 'hmac.New' returns a new HMAC hash using the given hash.Hash type and key.
	//            hash.Hash          key
	h := hmac.New(sha256.New, []byte(`Hz3Fe4578v3`))
	io.WriteString(h, data)

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {

	//Handle the root webpage
	http.HandleFunc("/", Homepage)
	//Handle the authentication page
	http.HandleFunc("/auth", auth)

	fmt.Println(http.ListenAndServe(":8080", nil))
}

var key string
func Homepage(w http.ResponseWriter, r *http.Request)  {

	//Check if there is a cookie for the 'session-id'
	cookie, err := r.Cookie("session-id")

	//If there is no cookie, create a new 'session-id' cookie
	if err != nil{
		cookie = &http.Cookie{
			Name: "session-id",
			Value: "",
			//Secure: true,
			HttpOnly: true,
		}
	}

	//Check if there was a key entered from a form value 'key'
	if r.FormValue("key") != ""{
		//Get the value of the POST 'key' and then hash it using 'getCode()'
		key = getCode(r.FormValue("key"))

		//Set the value of the 'session-id' cookie to the hashed key
		cookie.Value = key

		//Set the cookie to the browser
		http.SetCookie(w, cookie)
	}

	//Set the page headers to run the code below as HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Create a form for entering a key to hash
	io.WriteString(w, `<form method="POST">
	<label for="key">Enter a key</label>
	<input type ="text" name="key" id="key">
	<input type="submit">
	</br>
	<a href="/auth">Validate here</a>
	</form>
	</br>
	<h1>` + cookie.Value + `</h1>`)
}

//Authentication Page
func auth(w http.ResponseWriter, r *http.Request)  {

	//Set the headers to read text as 'text/html'
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Check if there is a 'session-id' cookie
	cookie, err := r.Cookie("session-id")

	//Check if the cookie exist by seeing if there is an error
	if err != nil{
		// if no cookie exist, redirect to the homepage
		http.Redirect(w, r, "/", 303)
		return
	}

	//Check if the cookie value is empty
	if cookie.Value == ""{
		// if cookie is empty, redirect to the homepage
		http.Redirect(w, r, "/", 303)
		return
	}

	//If the cookie.Value is the same as our stored key, then print the page
	if cookie.Value == key {
		io.WriteString(w, `<h1>` + cookie.Value + `</h1>` + `<h1>` + key + `</h1>`)
	}else{
		io.WriteString(w, `Key does not match`)
	}
}

