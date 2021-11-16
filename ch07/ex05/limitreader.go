package main

import (
	"bytes"
	"fmt"
	"io"
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

type LimReader struct {
	r     io.Reader
	i     int64 // 今の読んでるindex
	limit int64
}

func LimitReader(r io.Reader, n int64) *LimReader {
	return &LimReader{r, 0, n}
}

func (lm *LimReader) Read(p []byte) (n int, err error) {
	if (*lm).i >= (*lm).limit {
		return 0, io.EOF
	}
	fmt.Println(lm.i, lm.limit)
	n, err = lm.r.Read(p[lm.i:lm.limit])
	(*lm).i = int64(n)
	return n, err
}

func main() {
	a := StringsNewReader("ABCDE")
	fmt.Println(*a)
	b := LimitReader(a, 2)
	var byteBuffer bytes.Buffer
	io.Copy(&byteBuffer, b)
	fmt.Println(byteBuffer.Bytes())
}
