package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"basic-go/ch05/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(_url string) []string {
	fmt.Println(_url, "aa")

	sourceURL, err := url.Parse(_url)
	if err != nil {
		log.Print(err)
		return nil
	}

	saveDestination := filepath.Join("result", sourceURL.Host, sourceURL.Path)
	if strings.HasSuffix(sourceURL.Path, "/") || sourceURL.Path == "" {
		saveDestination = filepath.Join(saveDestination, "index.html")
	}

	err = os.MkdirAll(filepath.Dir(saveDestination), 0755)
	if err != nil {
		log.Print(err)
		return nil
	}

	file, err := os.Create(saveDestination)
	if err != nil {
		log.Print(err)
		return nil
	}

	resp, err := http.Get(_url)
	if err != nil {
		log.Print(err)
		return nil
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Print(err)
		return nil
	}
	resp.Body.Close()
	file.Close()

	list, err := links.Extract(_url)
	if err != nil {
		log.Print(err)
		return nil
	}
	var same_domain_links []string
	for _, l := range list {
		u, err := url.Parse(l)
		if err != nil {
			log.Print(err)
			return nil
		}
		if sourceURL.Hostname() == u.Hostname() {
			same_domain_links = append(same_domain_links, l)
		}
	}
	return same_domain_links
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
