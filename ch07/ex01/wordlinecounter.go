package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	var wc WordCounter
	var msg = "good morning\ndid u sleep well?"
	fmt.Fprint(&wc, msg)
	fmt.Println(wc)
	fmt.Fprint(&wc, msg)
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprint(&lc, msg)
	fmt.Println(lc)
	fmt.Fprint(&lc, msg)
	fmt.Println(lc)
}
