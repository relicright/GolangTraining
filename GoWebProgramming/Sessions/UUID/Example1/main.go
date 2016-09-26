package main

import (
	"net/http"
	"github.com/nu7hatch/gouuid"
	"fmt"
	"io"
)

func main() {

	http.HandleFunc("/", Homepage)

	fmt.Println(http.ListenAndServe(":8080", nil))
}

func Homepage(w http.ResponseWriter, r *http.Request)  {

	// Check if we have a cookie, if not cause an error to make one
	cookie, err := r.Cookie("session-id")

	// If no cookie was found create a new cookie
	if err != nil {

		//Create a new UUID (Universal Unique ID)
		id, _ := uuid.NewV4()

		//Create a new cookie to hold out UUID
		cookie = &http.Cookie{
			Name: "session-id",
			Value: id.String(),

			// Secure: can only be used on HTTPS
			// Secure: true,
			HttpOnly: true,
		}

		//Set the new cookie into the browser
		http.SetCookie(w, cookie)
	}

	fmt.Println(cookie)

	//Display the cookie value on the page
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<h1>` + cookie.Value + `</h1>`)

}
