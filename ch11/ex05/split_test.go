package split

import (
	"strings"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	s, sep := "a:b:c", ":"
// 	words := strings.Split(s, sep)
// 	if got, want := len(words), 3; got != want {
// 		t.Errorf("Split(%q, %q) returned %d words, want %d", s, sep, got, want)
// 	}
// }

func TestSplit(t *testing.T) {
	var tests = []struct {
		s         string
		sep       string
		wantedLen int
	}{
		{"a:b:c", ":", 3},
		{"dasd,dasd,w,rg,ttw", ",", 5},
		{"a:b:c", "@", 1},
	}
	for _, test := range tests {
		if got, want := len(strings.Split(test.s, test.sep)), test.wantedLen; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, want)
		}
	}
}
