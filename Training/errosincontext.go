package main

import (
	"fmt"
	"log"
	"math"
)

type NorgateMathError struct {
	lat, long string
	err error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt1(-5)
	if err != nil{
		log.Println(err)
	}
}

func sqrt1(f float64) (float64,error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"56.3", "25.6", nme}
	}

	return math.Sqrt(f), nil
}

