package main

import (
	"basic-go/ch08/ex10/links"
	"fmt"
	"log"
	"os"
)

func crawl(url string, done chan struct{}) []string {
	fmt.Println(url)
	list, err := links.Extract(url, done)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	done := make(chan struct{})

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, done)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
