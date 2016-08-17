package main

import "fmt"

func main() {
	newmap := map[string]string {
		"Zero": "first", "one": "second", "two": "third",
	}

	// Range loops over a slice or map and return a 'key/index' and a value
	     //key     value            map
	for key, value := range newmap {
		fmt.Println("The key/index is: ", key, " and the value of the key is:", value)
	}

	// It is important to note that a maps values and keys will be returned in a random fashion.
	// Map indexes will not be returned in the order they were entered into the map
	// if you run the code multiple times you will see that the map returns keys and values in
	// different orders every times.
}
