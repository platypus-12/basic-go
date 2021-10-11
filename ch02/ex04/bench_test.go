package main

import (
	"basic-go/ch02/ex04/bitshiftpopcount"
	"basic-go/ch02/ex04/popcount"
	"testing"
)

var result int

func BenchmarkPopCount(b *testing.B) {
	var r int
	for i := uint64(0); i < uint64(b.N); i++ {
		r += popcount.PopCount(i)
	}
	result = r
}

func BenchmarkBitShiftPopCount(b *testing.B) {
	var r int
	for i := uint64(0); i < uint64(b.N); i++ {
		r += bitshiftpopcount.BitShiftPopCount(i)
	}
	result = r
}
