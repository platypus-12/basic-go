package main

import "fmt"

func main() {
	var a, b []uint64
	a = append(a, 23, 45)
	b = append(b, 33)

	fmt.Println(a)
	fmt.Println(b)

	b = append(b, a...)
	fmt.Println(b)
}
