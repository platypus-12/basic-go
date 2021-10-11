package main

import (
	"fmt"
)

func main() {
	s := "abc"
	sb := &s
	fmt.Printf("%p\n", sb)
	fmt.Printf("%p\n", s[1:])
	fmt.Printf("%p\n", s[2:])
	b := []byte(s)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", b[1:])
	fmt.Printf("%p\n", b[2:])
	b[1] = 's'
	fmt.Printf("%p\n", b[1:])
	fmt.Printf("%v\n", b[1:])
	b = []byte("kkk")
	s2 := string(b)
	fmt.Printf("%p\n", b)
	fmt.Println(s2)
}
