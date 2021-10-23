package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v¥n", err)
		os.Exit(1)
	}
	hits := make(map[string]int)
	visit(doc, hits)
	for k, v := range hits {
		fmt.Println(k, v)
	}
}

func visit(n *html.Node, hits map[string]int) {
	if n.Type == html.ElementNode {
		hits[n.Data]++
	}

	if c := n.FirstChild; c != nil {
		visit(c, hits)
	}
	if c := n.NextSibling; c != nil {
		visit(c, hits)
	}
}
