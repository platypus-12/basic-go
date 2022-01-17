package main

import (
	"testing"
)

func TestAddAndHas(t *testing.T) {
	var a IntSet
	b := make(map[int]bool)
	nums := []int{1, 11111, 1111123, 324234, 4234}
	for _, n := range nums {
		a.Add(n)
		b[n] = true
	}
	var tests = []struct {
		want bool
		got bool
	}{
		{b[1], a.Has(1)},
		{b[3], a.Has(3)},
		{b[11111], a.Has(11111)},
		{b[1111123], a.Has(1111123)},
		{b[4234], a.Has(4234)},

	}
	for _, test := range tests {
		if test.want != test.got{
			t.Errorf("actual: %v want: %v\n", test.want, test.got)
		}
	}
}

func TestUnionWith(t *testing.T) {
	var a IntSet
	b := make(map[int]bool)
	nums := []int{1, 1123, 1111123, 324234, 4234}
	for _, n := range nums {
		a.Add(n)
		b[n] = true
	}
	




}