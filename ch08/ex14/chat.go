package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name string
	ch   chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			member := ""
			for client := range clients {
				member += client.name + ","
			}
			cli.ch <- "chat member here:" + member[:len(member)-1]
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	ch <- "Your name?"
	nameInput := bufio.NewScanner(conn)
	nameInput.Scan()
	who := nameInput.Text()
	// who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arricved"
	entering <- client{name: who, ch: ch}

	input := bufio.NewScanner(conn)

	timer := time.NewTimer(5 * time.Minute)
	defer timer.Stop()
	go func() {
		select {
		case <-timer.C:
			conn.Close()
		}
	}()
	for input.Scan() {
		timer.Reset(5 * time.Minute)
		messages <- who + ": " + input.Text()
	}

	leaving <- client{name: who, ch: ch}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
