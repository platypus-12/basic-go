package main

import (
	"fmt"
)

func main() {
	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)
	//ages["carol"] = 21  panic
	age, ok := ages["bob"]
	if !ok{
		fmt.Println(age)
	}
	fmt.Println(equal(map[string]int{"A":0}, map[string]int{"B":42}))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x{
		if yv, ok := y[k]; !ok || xv != yv{
			return false
		}
	}
	return true
}

