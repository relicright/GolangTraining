package main

import (

	//"io/ioutil"
	"io"
	"os"
	"log"
	"strings"
)

func main() {

	// Args hold the command-line arguments, starting with the program name.
	//when running this .go file in the command prompt, you can enter your
	//own string to use after using the  'go run filename.go' command.
	// I.E. - go run OSargs.go Barry
	name := os.Args[1]

	var str string = `<html>
	<head></head><body>
		Some words here ` + name + `
	</body>
	</html>`



	/*newStr := []byte(str)
	ioutil.WriteFile("./newindex.html", newStr, 0644)*/

	//************** or  *********************

	/*f, err := os.Create("./index1.html")
	if err != nil{
		log.Fatalln(err)
	}
	defer f.Close()
	io.WriteString(f, str)*/

	//***************** or ******************
	f, err := os.Create("./index2.html")
	if err != nil{
		log.Fatalln(err)
	}
	defer f.Close()

	_, err1 := io.Copy(f, strings.NewReader(str))
	if err1 != nil {
		log.Fatalln(err1)
	}
}