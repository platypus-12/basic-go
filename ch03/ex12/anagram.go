package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("asdfg", "sadfg"))
	fmt.Println(isAnagram("aaab", "abbb"))
	fmt.Println(isAnagram("日本語", "語日本"))
	fmt.Println(isAnagram("", ""))
}

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var rune1 = make([]rune, len(s1))
	var rune2 = make([]rune, len(s2))
	for i, v := range s1 {
		rune1[i] = v
	}
	for i, v := range s2 {
		rune2[i] = v
	}
	for _, v := range rune1 {
		position := locateSameRune(v, rune2)
		if position == -1 {
			return false
		}
		rune2 = append(rune2[:position], rune2[position+1:]...)
	}
	return true
}

func locateSameRune(r rune, l []rune) int {
	for i, v := range l {
		if r == v {
			return i
		}
	}
	return -1
}
