package main

import "fmt"

func main() {

	// Switch with an expression
	switch "Tim" {
	case "Tim":
		fmt.Printf("Hey Tim\n")
	case "Jenny":
		fmt.Printf("Hey Jenny\n")
	case "Marcus":
		fmt.Printf("Hey Marcus\n")
	case "Medhi":
		fmt.Printf("Hey Medhi\n")

		// Adding a fallthrough will act as if the next condition was automatically true
		// in this case: Julian's & Susan's 'case' would be evaluated as true
		fallthrough
	case "Julian":
		fmt.Printf("Hey Julian\n")
		fallthrough
	case "Susan":
		fmt.Printf("Hey Susan\n")

	// You can have multiple case evaluators in one 'case' conditions
	case "Eric", "Brandon":
		fmt.Printf("Hey Eric or Brandon")

	// You can have multiple case evaluators in one 'case' conditions
	case "Marvin", "Melvin":
		fmt.Printf("Both your names start with M")
	}

	// Switch with no expression
	myFriendsName := "Tim"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Hey my friend with a name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Hey Tim")
	case myFriendsName == "Tim", myFriendsName == "Medhi":
		fmt.Println("You are either Tim or Medhi")
	case myFriendsName == "Joe":
		fmt.Println("Hey Joe")
	}

	// Calls a switch function that evaluated the overload type
	SwitchOnType(Contact{})
}

// Switch can be on type

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an asset; asserting "x is of a certain type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}
