package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	url := os.Args[1]
	id := os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ouline2: %v¥n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing HTML: %s¥n", err)
		os.Exit(1)
	}
	node := ElementByID(doc, id)
	fmt.Println(node)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, findElement, findElement)
}

func findElement(n *html.Node, id string) bool {
	if n.Type != html.ElementNode {
		return false
	}
	for _, v := range n.Attr {
		if v.Key == "id" && v.Val == id {
			return true
		}
	}
	return false
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		hasFoundElement := pre(n, id)
		if hasFoundElement == true {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, id, pre, post)
		if node != nil {
			return node
		}
	}
	if post != nil {
		hasFoundElement := post(n, id)
		if hasFoundElement {
			return n
		}
	}
	return nil
}
