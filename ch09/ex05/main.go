package main

import (
	"fmt"
	"time"
)

func main() {
	chan1, chan2 := make(chan int), make(chan int)
	go func() {
		for x := range chan1 {
			x++
			fmt.Println(x)
			chan2 <- x

		}
	}()

	go func() {
		for y := range chan2 {
			y++
			fmt.Println(y)
			chan1 <- y
		}
	}()

	chan1 <- 0
	time.Sleep(1 * time.Second)
}

// 1回目 203091
// 2回目 195899
// 3回目 201120
