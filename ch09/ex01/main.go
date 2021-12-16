package main

import (
	"basic-go/ch09/ex01/bank1"
	"fmt"
	"time"
)

func main() {
	bank1.Deposit(2)
	go fmt.Println(bank1.Balance())
	go fmt.Println(bank1.Withdraw(2))
	go fmt.Println(bank1.Balance())
	go fmt.Println(bank1.Withdraw(2))
	time.Sleep(500 * time.Millisecond)
}
