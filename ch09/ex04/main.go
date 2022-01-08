package main

import (
	"fmt"
	"time"
)

func main() {
	startCh := make(chan int)
	var sendChan chan int
	recvChan := startCh

	for i := 0; i < 37874462; i++ {
		// fmt.Println(i)
		sendChan, recvChan = recvChan, make(chan int)
		go func(sendChan chan int, recvChan chan int) {
			tmp := <-sendChan
			recvChan <- tmp
		}(sendChan, recvChan)
	}
	now := time.Now()
	startCh <- 1
	<-recvChan
	fmt.Printf("経過: %vs\n", time.Since(now).Seconds())
}

// go run main.go
// 1
// 2
// 3
// 4
// 5
// .
// .
// .
// 37874461
// 37874462
// signal: killed

// % go run main.go
// 経過: 74.894298208s