package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 4, 5, 1}
	b := []int{1, 2, 3, 2, 1}
	c := []int{4, 5, 6, 7, 6, 5, 4}
	fmt.Println(sort.IntSlice(a))
	fmt.Println(IsPalindrome(sort.IntSlice(a)))
	fmt.Println(IsPalindrome(sort.IntSlice(b)))
	fmt.Println(IsPalindrome(sort.IntSlice(c)))
}
