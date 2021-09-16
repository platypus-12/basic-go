package main

import "fmt"

func main() {
	var p = f()
	fmt.Println(p)
	fmt.Println(f())
	fmt.Println(f())
}


func f() *int {
	v := 1
	return &v
}
