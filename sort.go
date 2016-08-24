package main

import (
	"fmt"
	"sort"
)

// A method can only be added to a type
type people []string

// The 'Sort()' function requires a sort interface.
// A sort interface has (3)three requirements:
// Len() int
// Swap()
// Less() bool
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {

	// Assign variable 'studygroup' the 'people' type
	studygroup := people{"Zeno", "John", "Al", "Jenny"}

	// 'sort.Sort()' requires a 'sort.interface' in order to sucessfully sort
	sort.Sort(studygroup)
	fmt.Println(studygroup)

	//////////////////////////////////////////////////////
	//// Sort through types without 'sort.interface'
	/////////////////////////////////////////////////////

	// In the instance that variable without type 'sort', but '[]string' needs to be sorted,
	// You would assign the '[]string' to a variable, in this case 's'
	s := []string{"Zeno", "John", "Al", "Jenny"}

	// You will then need to give 's' the 'sort.interface'
	// 'sort.StringSlice()' will attach the methods of  'sort.interface'
	// to []string
	newsort := sort.StringSlice(s)
	// or
	sort.Sort(newsort)
    // or 
    sort.StringSlice(s).Sort()

	// Now that 's' has the 'sort.interface' methods attached to it,
	// and it has been sorted above, we can print the result.
	fmt.Println(s)


    ////////////////////////////////////////////////////
    //// Sort through []int without 'sort.interface'
    ///////////////////////////////////////////////////

    n := []int {3,9,12,55,2,1,93}

    // 'sort.IntSlice()' attaches the methods of 'sort.interface' to
    // the variable in order for it to be compatible with the 'Sort()' function
    sort.IntSlice(n).Sort()

    // Print the sorted []int
    fmt.Println(n)
}
