package main

import (
	"basic-go/ch05/links"
	"flag"
	"fmt"
	"log"
)

type URL struct {
	path  string
	depth int
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	var depth = flag.Int("depth", 1, "depth")
	flag.Parse()

	worklist := make(chan []URL)
	unseenLinks := make(chan URL)

	// Start with the command-line arguments.
	go func() {
		var workURL []URL
		for _, x := range flag.Args() {
			workURL = append(workURL, URL{path: x, depth: 0})
		}
		worklist <- workURL
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for url := range unseenLinks {
				if url.depth > *depth {
					continue
				}
				foundLinks := crawl(url.path)
				go func() {
					var foundURLs []URL
					for _, x := range foundLinks {
						foundURLs = append(foundURLs, URL{path: x, depth: url.depth + 1})
					}
					worklist <- foundURLs
				}()
			}
		}()
	}

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.path] {
				seen[link.path] = true
				unseenLinks <- link
			}
		}
	}
}
