package main

import (
	"basic-go/ch02/ex03/looppopcount"
	"basic-go/ch02/ex03/popcount"
	"fmt"
)

func main() {
	fmt.Println(looppopcount.LoopPopCount(65535))
	fmt.Println(popcount.PopCount(65535))
	fmt.Println(looppopcount.LoopPopCount(256))
	fmt.Println(popcount.PopCount(256))
}
