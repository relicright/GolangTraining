package main

import (
"fmt"
)

func hello()  {
	fmt.Println("hello")
}

func world()  {
	fmt.Println("world")
}

func main() {

	// A more advanced defer that we want to run last
	defer retNum()

	// When you use the 'defer' keyword, the deferred expression is run at the end of the code block,
	// but also before the return
	defer world()
	hello()


}

func retNum()  {

	// Display to show the beginning of the block
	fmt.Println("Beginning of block")

	// You would think this would defer the expression to the very end of the block (which it will)
	defer fmt.Println("End of block")
	// But when you use two or more 'defer' keywords in the same block, the latest
	// deferred expression is put at the front of the list - in this case 'middle of block'
	// would print before 'end of block'
	defer fmt.Println("Middle of block")
}
