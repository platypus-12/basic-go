package main

import (
	"fmt"
)

func main() {
	medals := []string{"gold", "silver", "bronze"}
	//for i := len(medals) -1; i >= 0; i-- {
	for i := uint8(len(medals) -1); i >= 0; i-- {
		fmt.Println(medals[i])
	}
}
