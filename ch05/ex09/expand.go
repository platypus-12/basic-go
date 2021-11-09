package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("dasd $sWEs Ddasd $POsW $ああだsd", strings.ToLower))
	fmt.Println(expand("dasd $sWEs Ddasd $POsW $ああだsd", strings.ToUpper))
}

func expand(s string, f func(string) string) (str string) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if scanner.Text()[0] == 36 {
			str += fmt.Sprintf("%s ", f(scanner.Text()[1:]))
		} else {
			str += fmt.Sprintf("%s ", scanner.Text())
		}
	}
	return str
}
