package main

import (
	"fmt"
)

func main() {
	//type S []S
	var s []S
	s = append(s, s)
	fmt.Println(len(s))
//	fmt.Printf("%p %p %p \n", s, s[0], s[0][0])
	//fmt.Println(s)
}
