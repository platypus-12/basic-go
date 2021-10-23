package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sub(3, 1), math.Sin(2))
}

func sub(x, y int) (z int) {
	z = x - y
	return
}
