package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBR
	RMB
)


func main() {
	symbol := [...]string{USD: "$", EUR: "e", 1000: "g", RMB: "\\"}

	fmt.Println(RMB, symbol[RMB], symbol[100])
	fmt.Printf("%T", symbol)

	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1,2}
	fmt.Println(q)
	fmt.Println(r[2])
	
	k := [...]int{1,2,3}
	fmt.Printf("%T\n", k)


}
