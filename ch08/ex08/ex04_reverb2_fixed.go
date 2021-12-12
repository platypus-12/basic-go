package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	tick := time.Tick(1 * time.Second)

	chat := make(chan string)
	go func(chat chan<- string) {
		for input.Scan() {
			chat <- input.Text()
		}
	}(chat)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
		case str := <- chat:
			countdown = 10
			wg.Add(1)
			go echo(c, str, 1*time.Second, wg)
		}
	}
	wg.Wait()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
