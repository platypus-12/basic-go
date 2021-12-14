package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var done = make(chan struct{})
var response = make(chan string)

func mirroredQuery(urls []string) string {
	defer close(done)
	for _, url := range urls {
		req, _ := http.NewRequest("GET", url, nil)
		req.Cancel = done
		go func() {
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Print(err)
				return
			}
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				log.Print(err)
				return
			}
			response <- string(b)
		}()
	}
	s := <-response
	return s
}

// go run fetch.go http://gopl.io http://gopl.io http://gopl.io http://abehiroshi.la.coocan.jp/| less
func main() {
	urls := os.Args[1:]
	fmt.Println(mirroredQuery(urls))
}
