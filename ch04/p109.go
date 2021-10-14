package main

import (
	"fmt"
)

var m = make(map[string]int)

func main() {
	a := []string{"Hello", "World"}
	b := []string{"Hello", "世界"}
	fmt.Printf("%q\n", a[:1])
	fmt.Printf("%q\n", b[:])

	Add(a[:1])
	Add(a[:])
	Add(b[:1])
	fmt.Println(Count(a[:1]))
	m["aa"]++
	fmt.Println(m)
}

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
