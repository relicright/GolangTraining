package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"fmt"
)

func getCode(data string) string {

	// New returns a new HMAC hash using the given hash.Hash type and key.
	//            hash type        key
	h := hmac.New(sha256.New, []byte("ourkey"))

	//Write the passed in data the hash
	io.WriteString(h, data)

	//return the base 16 with lowercase a-f characters
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {

	// Create a code using the additional email data
	code := getCode("test@example.com")

	//Display the code in the console
	fmt.Println(code)
}
