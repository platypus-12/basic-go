package main

import (
	"basic-go/ch02/ex05/clearleastsignificantbitpopcount"
	"basic-go/ch02/ex05/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		popcount.PopCount(i)
	}
}

func BenchmarkClearLeastSignificantBitPopCount(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		clearleastsignificantbitpopcount.ClearLeastSignificantBitPopCount(i)
	}
}
