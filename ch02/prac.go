package main

import (
		"fmt"
		"basic-go/ch02/popcount"
)

func main() {
	fmt.Printf("%v\n", 1&2)
	fmt.Printf("%v\n",5&6)

	fmt.Printf("%v\n", 12&11)

	fmt.Printf("%v\n", 12&12)
	fmt.Printf("%v\n", popcount.PopCount(4))
	fmt.Printf("")

}
