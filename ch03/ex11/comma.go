package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-123"))
	fmt.Println(comma("1234.56"))
	fmt.Println(comma("123.456789"))
	fmt.Println(comma("+123.456789"))
	fmt.Println(comma("+1234567.89"))
	fmt.Println(comma("-12345.6789"))
}

func comma(s string) string {
	if s[0] == '+' || s[0] == '-' {
		return string(s[0]) + comma(s[1:])
	}
	n := len(s)
	if 3 >= n {
		return s
	}
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		return comma(s[:dot]) + s[dot:]
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
