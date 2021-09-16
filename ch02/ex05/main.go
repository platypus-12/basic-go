package main

import (
	"basic-go/ch02/ex05/clearleastsignificantbitpopcount"
	"basic-go/ch02/ex05/popcount"
	"fmt"
)

func main() {
	fmt.Println(clearleastsignificantbitpopcount.ClearLeastSignificantBitPopCount(65535))
	fmt.Println(popcount.PopCount(65535))
	fmt.Println(clearleastsignificantbitpopcount.ClearLeastSignificantBitPopCount(256))
	fmt.Println(popcount.PopCount(256))
}
