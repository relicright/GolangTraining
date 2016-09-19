package main
import (
  "fmt"
)

func main()  {
  var p *int

  if p != nil {
    //THE '*' REFERENCES THE VALUE THE POINTER IS POINTING AT
    fmt.Println("Value of P:", *p)
  }else{
    fmt.Println("P is nil")
  }

  var v int = 42

  // THE '&' MEANS CONNECT THE POINTER TO THIS VARIABLE
  p = &v

  if p != nil {
    //THE '*' REFERENCES THE VALUE THE POINTER IS POINTING AT
    fmt.Println("Value of P:", *p)
  }else{
    fmt.Println("P is nil")
  }

  var value1 float64 = 42.13

  pointer1 := &value1
  fmt.Println("Value 1:", *pointer1)

  *pointer1 = *pointer1 / 31
  fmt.Println("value1 is:", *pointer1)
  fmt.Println("value1 is:", value1)

}
