package main

import (
	"basic-go/ch02/ex03/looppopcount"
	"basic-go/ch02/ex03/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		popcount.PopCount(i)
	}
}

func BenchmarkLoopPopCount(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		looppopcount.LoopPopCount(i)
	}
}
