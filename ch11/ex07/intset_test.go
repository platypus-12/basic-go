package intset

import (
	"math/rand"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	var intSet IntSet
	for i := 0; i < b.N; i++ {
		r := rand.Intn(10000000) //うまくAddだけのベンチマークが測れない
		intSet.Add(r)
	}
}

var s1 IntSet
var s2 IntSet

func BenchmarkUnionWith(b *testing.B) {
	s1.Add(1)
	for i := 0; i < b.N; i++ {
		s2.Add(rand.Intn(10000000)) //うまくUnionWithだけのベンチマークが測れない
		s1.UnionWith(&s2)
	}

}

// words []uint64のとき
// % go test -bench=.
// goos: darwin
// goarch: arm64
// pkg: basic-go/ch11/ex07
// BenchmarkAdd-8          87104223                13.61 ns/op
// BenchmarkUnionWith-8       10000            104200 ns/op
// PASS
// ok   basic-go/ch11/ex07 3.239s


