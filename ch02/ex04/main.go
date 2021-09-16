package main

import (
	"basic-go/ch02/ex04/bitshiftpopcount"
	"basic-go/ch02/ex04/popcount"
	"fmt"
)

func main() {
	fmt.Println(bitshiftpopcount.BitShiftPopCount(65535))
	fmt.Println(popcount.PopCount(65535))
	fmt.Println(bitshiftpopcount.BitShiftPopCount(256))
	fmt.Println(popcount.PopCount(256))
}
