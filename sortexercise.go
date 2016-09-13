package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type UserRequst struct {
	Name      string
	os        string
	Response  io.Reader
	ResString string
}

func (u UserRequst) UserGet(s string) io.Reader {

	res, err := http.Get(s)

	if err != nil {
		log.Fatalln(err)
	}

	return res.Body
}

func (u UserRequst) ResponseToString(r io.Reader) string {

	bs, _ := ioutil.ReadAll(r)
	return string(bs)
}

func main() {

	u := UserRequst{"James", "Bond", nil, ""}
	u.Response = u.UserGet("http://www.wellnesscenters4life.com")
	u.ResString = u.ResponseToString(u.Response)
	//u.os = os.Getenv()

	fmt.Println(u.ResString)
}
