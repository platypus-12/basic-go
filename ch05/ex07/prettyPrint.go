package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ouline2: %v¥n", err)
			continue
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "parsing HTML: %s¥n", err)
			continue
		}
		forEachNode(doc, startElement, endElement)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, v := range n.Attr {
			fmt.Printf(" %s=\"%s\"", v.Key, v.Val)
		}
		fmt.Printf(">\n")
		depth++
		return
	}
	if n.Type == html.ElementNode && n.FirstChild == nil {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, v := range n.Attr {
			fmt.Printf(" %s=\"%s\"", v.Key, v.Val)
		}
		fmt.Printf("/>\n")
		return
	}
	if n.Type == html.CommentNode {
		fmt.Println("<!--", n.Data, "-->")
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
		return
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil && n.FirstChild != nil {
		post(n)
	}
}
