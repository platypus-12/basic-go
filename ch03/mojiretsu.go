package main

import (
	"fmt"
)

func main(){
	s := "hello, world"
	d := "あいうえお"
  fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(len(d))
	fmt.Println(d[1], d[3])

	fmt.Println(s[0:5])
	fmt.Println(s[0:0])

	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])
	k := `ddasd
	dasd
	asdasdasd
	 fadfafasf

	 fasffsa
	 Faf`
	fmt.Println(k)
}
