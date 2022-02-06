package main

import "basic-go/ch12/ex02/display"


func main() {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	display.Display("c", c)
}
