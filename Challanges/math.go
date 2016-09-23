package main

import(
  "fmt"
  "math/big"
  "math"
)

func main()  {
  fmt.Println(mathofsametype(22.2, 33.4))

// BIG.FLOAT ALLOWS FOR MORE ACCURATE CALCULATIONS
  var b1, b2, b3, bigSum big.Float

  //BIG.FLOAT VARIABLES ARE SET THROUGH DOT SYNTEX - IN THIS CASE .SETFLOAT64()
  b1.SetFloat64(23.5)
  b2.SetFloat64(33.5)
  b3.SetFloat64(28.2)


  //BIG FLOATS MATH OPERATIONS ARE THROUGH THE DOT SYNTEX - IN THIS CASE .ADD()
  bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
  fmt.Printf("Bigsum equlas: %10g\n", &bigSum)

  circleRadius := 15.5
  circumference := circleRadius * math.Pi

  fmt.Printf("Circumference Equals: %.2f\n", circumference)
}

// VARIABLES MUST BE OF THE SAME TYPE TO DO MATH OPERATIONS ON THEM
func mathofsametype(x, y float64) float64{
  return x + y
}
