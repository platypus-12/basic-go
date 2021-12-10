package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		placeInfo := strings.Split(arg, "=")
		go displayClock(placeInfo)
	}
	for {
	}
}

func displayClock(placeInfo []string) {
	conn, err := net.Dial("tcp", placeInfo[1])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		time, _ := reader.ReadString('\n')
		fmt.Printf("%s  %s\n", placeInfo[0], time[:len(time)-1])
	}
}
