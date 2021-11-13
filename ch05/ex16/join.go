package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join(",", "aaa", "pdas"))
}

//func Join(str []string, sep string) string
func join(sep string, str ...string) string {
	return strings.Join(str, sep)
}
