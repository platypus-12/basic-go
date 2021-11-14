package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	// defer trace("bigSlowOperation")()
	defer trace1("bigSlowOperation")
	time.Sleep(10 * time.Second)
	fmt.Println("fin")
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func trace1(msg string) {
	start := time.Now()
	log.Printf("enter %s", msg)
	log.Printf("exit %s (%s)", msg, time.Since(start))
	return
}

func main() {
	double(2)
	fmt.Println(triple(3))
	bigSlowOperation()
}

func _double(x int) int {
	return x + x
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
