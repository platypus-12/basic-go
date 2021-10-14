package main

import (
	"fmt"
)

func main() {
	var x []int
  x = appendInt(x, 1)
	x = appendInt(x, 2, 3)
	x = appendInt(x, 4,5,6)
	x = appendInt(x, x...)
	fmt.Println(x)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
