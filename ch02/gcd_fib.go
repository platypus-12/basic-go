package main

import (
	"fmt"
)

func main() {
	fmt.Println(euclideanGcd(8, 12))
	fmt.Println(euclideanGcd(12, 8))
	fmt.Println(fib(5))
}

func euclideanGcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
