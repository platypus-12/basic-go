package main

import (
	"basic-go/ch02/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
