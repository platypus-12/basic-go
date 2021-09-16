package main

import (
	"fmt"
	"basic-go/ch02/ex01/tempconv"
)

func main() {
	fmt.Println(tempconv.CToK(tempconv.BoilingC))
	fmt.Println(tempconv.KToC(tempconv.CToK(tempconv.BoilingC)))
}
