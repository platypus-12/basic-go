package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	url := "https://golang.org"
	// url := os.Args[1]
	// id := os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline2: %v¥n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing HTML: %s¥n", err)
		os.Exit(1)
	}
	elementNodes := ElementByTagName(doc, "h1", "h2", "h3", "img")
	for _, node := range elementNodes {
		fmt.Println(node.Data)
	}
}

func ElementByTagName(doc *html.Node, tagNames ...string) []*html.Node {
	return forEachNode(doc, tagNames, findTagName, nil)
}

func findTagName(n *html.Node, tagnames []string) bool {
	for _, tagname := range tagnames {
		if n.Type == html.ElementNode && n.Data == tagname {
			return true
		}
	}
	return false
}

func forEachNode(n *html.Node, tagNames []string, pre, post func(n *html.Node, tagnames []string) bool) (nodes []*html.Node) {
	if pre != nil {
		hasFoundElement := pre(n, tagNames)
		if hasFoundElement == true {
			nodes = append(nodes, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, tagNames, pre, post)
		if node != nil {
			nodes = append(nodes, node...)
		}
	}
	if post != nil {
		hasFoundElement := post(n, tagNames)
		if hasFoundElement {
			nodes = append(nodes, n)
		}
	}
	return nodes
}
