package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input         string
		wantedCounts  map[rune]int
		wantedUtflen  [utf8.UTFMax + 1]int
		wantedInvalid int
	}{
		{"aあaaaa", map[rune]int{'a': 5, 'あ': 1}, [utf8.UTFMax + 1]int{0, 5, 0, 1, 0}, 0},
		{"", map[rune]int{}, [utf8.UTFMax + 1]int{0, 0, 0, 0, 0}, 0},
		{"aαあ😀", map[rune]int{'a': 1, 'α': 1, 'あ': 1, '😀': 1}, [utf8.UTFMax + 1]int{0, 1, 1, 1, 1}, 0},
	}
	for _, test := range tests {
		in := bufio.NewReader(strings.NewReader(test.input))
		gotCounts, gotUtflen, gotInvalid := charCount(in)
		if !reflect.DeepEqual(gotCounts, test.wantedCounts) {
			t.Errorf("counts\ninput: %q\nactual: %v\n  want: %v", test.input, gotCounts, test.wantedCounts)
		}
		if gotUtflen != test.wantedUtflen {
			t.Errorf("utflen\ninput: %q\nactual: %v\n  want: %v", test.input, gotUtflen, test.wantedUtflen)
		}
		if gotInvalid != test.wantedInvalid {
			t.Errorf("invalid\ninput: %q\nactual: %v\n  want: %v", test.input, gotInvalid, test.wantedInvalid)
		}
	}
}
