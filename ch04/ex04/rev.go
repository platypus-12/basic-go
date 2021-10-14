package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 4)
	fmt.Println(a)
}

func rotate(s []int, r int) {
	if len(s) == 0 || r == 0 {
		return
	}
	var queue []int
	for i, _ := range s {
		if i < r {
			queue = append(queue, s[i])
		}
		if len(s)-r <= i {
			s[i] = queue[0]
			queue = queue[1:]
			continue
		}
		if i+r < len(s) {
			s[i] = s[i+r]
		}
	}
}
