package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

//https://cs.opensource.google/go/go/+/refs/tags/go1.17.3:src/io/io.go;l=455
//https://cs.opensource.google/go/go/+/refs/tags/go1.17.3:src/strings/reader.go;l=160

type StringsReader struct {
	s string
	i int64 // 今の読んでるindex
}

func (sr *StringsReader) Read(p []byte) (n int, err error) {
	if (*sr).i >= int64(len((*sr).s)) {
		return 0, io.EOF
	}
	n = copy(p, sr.s[(*sr).i:])
	(*sr).i += int64(n)
	return n, err
}

func StringsNewReader(s string) *StringsReader {
	return &StringsReader{s, 0}
}

func main() {
	doc, err := html.Parse(strings.NewReader("<title>test</title>"))
	fmt.Println(doc, err)
	fmt.Println(doc.Type)
	doc1, err1 := html.Parse(StringsNewReader("<title>test</title>"))
	fmt.Println(doc1, err1)
	fmt.Println(doc1.Type)
}
