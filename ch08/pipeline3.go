package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go square(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int){
	for x := range in{
		fmt.Println(x)
	}
}