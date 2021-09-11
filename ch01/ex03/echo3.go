package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.5fs elapsed()\n", time.Since(start1).Seconds())

	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.5fs elapsed\n", time.Since(start2).Seconds())
}
