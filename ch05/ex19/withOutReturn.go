package main

import "fmt"

func main() {
	fmt.Println(withOutReturn())
}

func withOutReturn() (a int) {
	defer func() {
		_ = recover()
		a = 1
	}()
	panic(1)
}
