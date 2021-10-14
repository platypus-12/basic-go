package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(different_bit(c1, c2))
	c3 := [32]uint8{1, 1, 1, 1, 1}
	c4 := [32]uint8{1, 0, 1, 1, 1}
	fmt.Println(different_bit(c3, c4))

}

func different_bit(a [32]uint8, b [32]uint8) uint8 {
	var _sum uint8
	for i := 0; i < len(a); i++ {
		_sum += (pc[a[i]^b[i]])
	}
	return _sum
}
