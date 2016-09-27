package main

import (
	"net/http"
	"io"
)

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request)  {

	//Check if a cookie exist with the name 'logged-in'
	cookie, err := r.Cookie("logged-in")

	// If no cookie exist, create one
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: "logged-in",
			Value: "0",
			//Secure: true,
			HttpOnly: true,
		}
	}

	//Check log in
	if r.Method == "POST"{
		password := r.FormValue("password")
		if password == "secret"{
			cookie = &http.Cookie{
				Name: "logged-in",
				Value: "1",
				//Secure: true,
				HttpOnly: true,
			}
		}
	}

	//if logout, then logout
	if r.URL.Path == "/logout"{
		cookie = &http.Cookie{
			Name: "logged-in",
			Value: "0",
			MaxAge: -1,
			//Secure: true,
			HttpOnly: true,
		}
	}

	//Set the cookie after determining what state it should be returned as
	http.SetCookie(w, cookie)

	//Create a string variable to hold the page HTML in
	var html string

	// not logged in page
	if cookie.Value == "0"{
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1>LOG IN</h1>
			<form method="POST">
				<h3>Password</h3>
				<input type="text" name="password">
				<br>
				<input type="submit">
			</form>
			</body>
			</html>`
	}

	// logged in page
	if cookie.Value == "1" {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1><a href="/logout">LOG OUT</a></h1>
			</body>
			</html>`
	}

	//Write the page to the writer
	io.WriteString(w, html)
}


// NOT GOOD PRACTICE
// adding user data to a cookie
// with no way of knowing whether or not
// they might have altered that data
//
// HMAC would allow us to determine
// whether or not the data in the cookie was altered
//
// however, best to store user data
// on the server
// and keep backups
