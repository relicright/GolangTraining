package main

import (
	"fmt"
	"path/filepath"
	"log"
	"os"
)

func main() {

	// Create a variable that will contain the filepath
	// Abs returns an absolute representation of path.
	// If the path is not absolute it will be joined with the current
	// working directory to turn it into an absolute path.
	root, _ := filepath.Abs(".")
	fmt.Println("Processing path", root)

	//
	// Walk walks the file tree rooted at root, calling walkFn for each file or
	// directory in the tree, including root. All errors that arise visiting files
	// and directories are filtered by walkFn.
	err := filepath.Walk(root, processPath)
	if err != nil{
		log.Fatal(err)
	}
}

// Now we create the function that will be called for each file and directory as
// it is being walked through
func processPath(path string, info os.FileInfo, err error) error  {

	// Check for errors
	if err != nil {
		return err
	}

	// Check if the path is a root
	if path != "."{
		// If the path is a root then it is a directory
		if info.IsDir(){
			fmt.Println("directory:", path)
		}else{
			// Otherwise the path is a file
			fmt.Println("File:", path)
		}
	}

	// Return an error object that is 'nil'
	return nil
}
