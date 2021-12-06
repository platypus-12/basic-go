package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// var ErrNotExist = errors.New("file does not exist")
	_, err := os.Open("/no/such/file")
	fmt.Println(IsNotExist(err))
}

func IsNotExist(err error) bool {
	fmt.Printf("%#v\n", err)
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}
	fmt.Printf("%#v\n", err)
	return err == syscall.ENOENT || err == os.ErrNotExist
}
