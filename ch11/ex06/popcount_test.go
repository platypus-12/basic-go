package popcount

import (
	"basic-go/ch02/ex04/bitshiftpopcount"
	"basic-go/ch02/ex05/clearleastsignificantbitpopcount"
	"basic-go/ch02/popcount"
	"testing"
)

func benchmark(b *testing.B, size uint64, f func(x uint64) int) {
	for i := 0; i < b.N; i++ {
		f(size)
	}
}

func benchmarkInit(b *testing.B, size uint64, f func(x uint64) int) {
	var pc [256]byte
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	for i := 0; i < b.N; i++ {
		f(size)
	}
}

func BenchmarkPoPCount1(b *testing.B)     { benchmarkInit(b, 1, popcount.PopCount) }
func BenchmarkPoPCount10(b *testing.B)    { benchmarkInit(b, 10, popcount.PopCount) }
func BenchmarkPoPCount100(b *testing.B)   { benchmarkInit(b, 100, popcount.PopCount) }
func BenchmarkPoPCount1000(b *testing.B)  { benchmarkInit(b, 1000, popcount.PopCount) }
func BenchmarkPoPCount10000(b *testing.B) { benchmarkInit(b, 10000, popcount.PopCount) }

func BenchmarkBitShift1(b *testing.B)     { benchmark(b, 1, bitshiftpopcount.BitShiftPopCount) }
func BenchmarkBitShift10(b *testing.B)    { benchmark(b, 10, bitshiftpopcount.BitShiftPopCount) }
func BenchmarkBitShift100(b *testing.B)   { benchmark(b, 100, bitshiftpopcount.BitShiftPopCount) }
func BenchmarkBitShift1000(b *testing.B)  { benchmark(b, 1000, bitshiftpopcount.BitShiftPopCount) }
func BenchmarkBitShift10000(b *testing.B) { benchmark(b, 10000, bitshiftpopcount.BitShiftPopCount) }

func BenchmarkClearLeastSignificantBit1(b *testing.B) {
	benchmark(b, 1, clearleastsignificantbitpopcount.ClearLeastSignificantBitPopCount)
}
func BenchmarkClearLeastSignificantBit10(b *testing.B) {
	benchmark(b, 10, clearleastsignificantbitpopcount.ClearLeastSignificantBitPopCount)
}
func BenchmarkClearLeastSignificantBit100(b *testing.B) {
	benchmark(b, 100, clearleastsignificantbitpopcount.ClearLeastSignificantBitPopCount)
}
func BenchmarkClearLeastSignificantBit1000(b *testing.B) {
	benchmark(b, 1000, clearleastsignificantbitpopcount.ClearLeastSignificantBitPopCount)
}
func BenchmarkClearLeastSignificantBit10000(b *testing.B) {
	benchmark(b, 10000, clearleastsignificantbitpopcount.ClearLeastSignificantBitPopCount)
}

// ex06 % go test -bench=.
// goos: darwin
// goarch: arm64
// pkg: basic-go/ch11/ex06
// BenchmarkPoPCount1-8                            554154667                2.034 ns/op
// BenchmarkPoPCount10-8                           586188066                2.028 ns/op
// BenchmarkPoPCount100-8                          596728311                2.023 ns/op
// BenchmarkPoPCount1000-8                         596632136                2.033 ns/op
// BenchmarkPoPCount10000-8                        592266357                2.039 ns/op
// BenchmarkBitShift1-8                            55608223                21.23 ns/op
// BenchmarkBitShift10-8                           56182405                21.17 ns/op
// BenchmarkBitShift100-8                          55934493                21.18 ns/op
// BenchmarkBitShift1000-8                         55779362                21.12 ns/op
// BenchmarkBitShift10000-8                        55592336                21.12 ns/op
// BenchmarkClearLeastSignificantBit1-8            1000000000               0.9319 ns/op
// BenchmarkClearLeastSignificantBit10-8           638631768                1.868 ns/op
// BenchmarkClearLeastSignificantBit100-8          550880298                2.174 ns/op
// BenchmarkClearLeastSignificantBit1000-8         385737205                3.124 ns/op
// BenchmarkClearLeastSignificantBit10000-8        425545118                2.816 ns/op
