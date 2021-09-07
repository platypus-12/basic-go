package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := ""
	for i, arg := range os.Args[:] {
		s += strconv.Itoa(i) + " " + arg + "\n"
	}
	fmt.Print(s)
}
