package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func del_duplicated(bytes []byte) []byte {
	var tmp rune
	j := 0
	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		fmt.Println(i, r, size)
		if unicode.IsSpace(tmp) == false || unicode.IsSpace(r) == false {
			copy(bytes[j:j+size], bytes[i:i+size])
			j += size
		}
		i += size
		tmp = r
	}
	return bytes[:j]

}

func main() {
	data := "sample  ああ 尾   sssaaa   daw"
	data_byte := []byte(data)
	fmt.Printf("%s\n", del_duplicated(data_byte))
}
