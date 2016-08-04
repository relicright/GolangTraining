package main

import (
	"fmt"
	// Fully qualified path of package to import
	"github.com/relicright/GolangTraining/stringutils"
)

func main() {
	// Using an exposed method in the reverse.go file
	fmt.Println(stringutils.Reverse("test string"))
}
