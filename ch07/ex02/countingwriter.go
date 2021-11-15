package main

import (
	"fmt"
	"io"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WriteWrapper struct {
	writer   io.Writer
	bytesNum int64
}

func (ww *WriteWrapper) Write(p []byte) (int, error) {
	n, err := (*ww).writer.Write(p)
	(*ww).bytesNum = int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var wwptr WriteWrapper
	wwptr.writer = w
	return &wwptr, &wwptr.bytesNum
}

func main() {
	var c ByteCounter
	ww, n := CountingWriter(&c)
	var name = "Dolly"
	fmt.Println(*n)
	fmt.Fprintf(ww, "hello, %s", name)
	fmt.Println(*n)
}
