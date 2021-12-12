package main

import (
	"basic-go/ch05/links"
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0 ; n-- {
		list := <- worklist
		for _, link := range list {
			if !seen[link] {
				n++
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
