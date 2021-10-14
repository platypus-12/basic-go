package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	_sha384 := flag.Bool("sha384", false, "enable sha384")
	_sha512 := flag.Bool("sha512", false, "enable sha512")
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if *_sha384 == false && *_sha512 == false {
			fmt.Printf("%x\n", sha256.Sum256([]byte(input.Text())))
			continue
		}
		if *_sha384 == true {
			fmt.Printf("%x\n", sha512.Sum384([]byte(input.Text())))
		}
		if *_sha512 == true {
			fmt.Printf("%x\n", sha512.Sum512([]byte(input.Text())))
		}
	}
}
