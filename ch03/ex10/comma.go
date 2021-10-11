package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1234567891234"))
}

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		fmt.Fprintf(&buf, "%c", s[i])
		if i != len(s)-1 && (len(s)-i)%3 == 1 {
			buf.WriteByte(',')
		}
	}

	return buf.String()
}
