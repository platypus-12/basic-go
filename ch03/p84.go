package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))
	s := fmt.Sprintf("x=%d", x)
	fmt.Println(s)
	a, _ := strconv.Atoi("123")
	b, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(a, b)
}


