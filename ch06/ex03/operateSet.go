package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)

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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
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
		for j := 0; j < 64; j++ {
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
	word, bit := x/64, uint(x%64)
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
	t := make([]uint64, len(s.words))
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
	len(s)

}

func main() {
	var x IntSet
	x.AddAll(1, 3, 4, 144)
	fmt.Println(x.String())
}
