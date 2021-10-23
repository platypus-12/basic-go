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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	if c := n.FirstChild; c != nil && !(c.Data == "script" || c.Data == "style") {
		links = visit(links, c)
	}
	if c := n.NextSibling; c != nil && !(c.Data == "script" || c.Data == "style") {
		links = visit(links, c)
	}
	return links
}
