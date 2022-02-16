package main

import (
	"basic-go/ch13/ex01/equal"
	"fmt"
)

func main() {
	fmt.Println(equal.Equal(1.0, 1.0+1e-9))
	fmt.Println(equal.Equal(1.0, 1.0+1e-10-1e-9))
	fmt.Println(equal.Equal(1.0i, 1.0i+1e-9))
	fmt.Println(equal.Equal(1.0i, 1.0i+1e-10-1e-9))
}
