package main

import (
	"basic-go/ch13/ex02/equal"
	"fmt"
)


func main() {
	type CyclePtr *CyclePtr
	var cyclePtr1 CyclePtr
	cyclePtr1 = &cyclePtr1

	fmt.Println(equal.IsCycle(cyclePtr1))
	fmt.Println(equal.IsCycle("cycle"))
}
