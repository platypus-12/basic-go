package main

import (
	"fmt"
)

func main() {
	// ch := make(chan int, 4) buffer が多いとコインを投げる
	ch := make(chan int, 4)

	for i := 0; i < 10; i++ {
		fmt.Println("cnt" , i)
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
