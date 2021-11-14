package main

import (
	"bytes"
	"fmt"
)

const cpuRegisterSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/cpuRegisterSize, uint(x%cpuRegisterSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/cpuRegisterSize, uint(x%cpuRegisterSize)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < cpuRegisterSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", cpuRegisterSize*i+j)
			}

		}

	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() (length int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < cpuRegisterSize; j++ {
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}
	return
}

func (s *IntSet) Remove(x int) {
	if len(s.words) == 0 {
		return
	}
	word, bit := x/cpuRegisterSize, uint(x%cpuRegisterSize)
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	if len(s.words) == 0 {
		return nil
	}
	var _copy IntSet
	t := make([]uint, len(s.words))
	_ = copy(t, s.words)
	_copy.words = t
	return &_copy
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elems() (elems []uint) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < cpuRegisterSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, uint(cpuRegisterSize*i+j))
			}
		}
	}
	return elems
}

func main() {
	var x, y IntSet
	x.AddAll(1, 3, 4, 144)
	y.AddAll(3, 144, 3, 5, 3455345)
	fmt.Println(&x)
	fmt.Println(x.Has(4))
	x.Add(111)
	fmt.Println(&x)
	fmt.Println(x.Len())
	x.Remove(144)
	fmt.Println(&x)
	for i, v := range x.Elems() {
		fmt.Println(i, v)
	}
}
