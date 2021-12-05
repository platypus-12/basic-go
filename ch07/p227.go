package main

import (
	"fmt"
	"syscall"
)

func main() {
	fmt.Printf("%.7g\n", 167.00011673013586)
	var uu uintptr = uintptr(2)
	var tt = int(3)
	fmt.Println(tt)
	fmt.Println(uu)
	var err error = syscall.Errno(2)
	fmt.Println(syscall.Errno(2))
	fmt.Println(err)
	fmt.Println(err.Error())
	fmt.Println(err)
}
