package main

import (
	"basic-go/ch02/ex04/bitshiftpopcount"
	"basic-go/ch02/ex04/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		popcount.PopCount(i)
	}
}

func BenchmarkBitShiftPopCount(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		bitshiftpopcount.BitShiftPopCount(i)
	}
}
