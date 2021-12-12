package main

import (
	"basic-go/ch05/links"
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 2; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {worklist <- foundLinks}()
				// 質問 p280 ここをgoroutine にしている理由
				// どのようにデッドロックが起きる？
				// worklist <- foundLinks
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
