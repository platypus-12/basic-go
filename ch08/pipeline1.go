package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			if x == 1000000 {
				break
			}
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
	}
}
