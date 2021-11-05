package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v¥n", err)
			continue
		}
		fmt.Println(url)
		fmt.Println("images:", images, "words", words)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "svg") {
		images++
	}

	if n.Type == html.TextNode {
		words = wordFreq(n.Data)
	}

	for c := n.FirstChild; c != nil && !(c.Data == "script" || c.Data == "style"); c = c.NextSibling {
		word, image := countWordsAndImages(c)
		words, images = words+word, images+image
	}
	return
}

func wordFreq(s string) (word int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word++
	}
	return
}
