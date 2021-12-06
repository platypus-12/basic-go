package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	// c := w.(*bytes.Buffer) もってないのでぱにっく
	fmt.Println(f)

	// rw := w.(io.ReadWriter)
	// w = new(ByteCounter)
	// rw = w.(io.ReadWriter)
	// fmt.Println(rw)

	w = os.Stdout
	//2つの結果を待ち受けるとpanic
	ff, ok := w.(*os.File)
	fmt.Println(ff, ok)
	bb, ok := w.(*bytes.Buffer)
	fmt.Println(bb, ok)

	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
}
