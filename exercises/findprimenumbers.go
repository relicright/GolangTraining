package main

import "fmt"

func main() {

	// A function expression assigns the function (func()) to the variable to be used
	// any other place withing the 'main' function scope
	expression1, iseven := checkfunc(4)

	/*expression2, even := func (n int) (int, bool) {

		return n/2, n%2 == 0
	}*/

	if iseven == true {
		fmt.Println("it is even: ", expression1)
	}


	fmt.Printf("%T", expression1)

	fmt.Println(expression1)

	numtotest := []int {22, 33, 54, 645, 33}
	println("Variadic num: ", variadicpara(numtotest...))

	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,2,3,4}
	foo(aSlice...)
	foo()

	//findtheprime(600851475143)
}

func checkfunc(n int) (int, bool) {
	return n/2, n%2 == 0
}

func variadicpara(nums ...int ) int  {

	var largestnum int
	for _, rng := range nums {
		if rng > largestnum {
			largestnum = rng
		}
	}
	return largestnum
}

func foo(n ...int) int {

	var x int
	for _, val := range n {
		if val > 1 {
			fmt.Println("greater than one")
		}else{
			fmt.Println("less than one")
		}
	}

	return x
}

func findtheprime(n int)  {

	var primes []int

	for i:=1; i < n; i++ {

		if n % i == 0 {
			primes = append(primes, i)
			fmt.Println(primes)
		}
	}

	fmt.Println(primes)
}
