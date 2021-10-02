package main

import (
	"fmt"
	"math"
)

func main(){
	var f float32 = 16777216
	fmt.Println(f == f + 1)
  const e = 2.71828
	const Avogadro = 6.02214129e23
  const Plank = 6.62606957e-34

	for x:= 0; x < 8; x++ {
		fmt.Printf("x = %d  e^x = %8.3f\n", x, math.Exp(float64(x)))
		fmt.Printf("x = %d  e^x = %8.3e\n", x, math.Exp(float64(x)))
		fmt.Printf("x = %d  e^x = %1.4g\n", x, math.Exp(float64(x)))
	}
}
