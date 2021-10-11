package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	ss := "プログラムa"
	fmt.Printf("% x\n", ss)
	r := []rune(ss)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
	fmt.Println(string(65))
	fmt.Println(string(0x41))
	fmt.Println(string(0x4eac))
	fmt.Println(string(0x30d7))
	fmt.Println(string(1234556))

	fmt.Println("-------")
	n := 0
	for _, _ = range "ajdapsfあああ" {
		n++
	}
	fmt.Println(n)

	m := 0
	for range "dadasdいい" {
		m++
	}
	fmt.Println(m)

	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i, r := range "Hello, 世界aA" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println("+++++++++++")
	fmt.Println(utf8.DecodeRuneInString("Hsjdasd"))
	fmt.Println("+++++++++++")

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	s1 := "qwertyuiop"
	fmt.Println(HasPrefix(s1, "qwe"), "true expected")
	fmt.Println(HasPrefix(s1, "aaa"), "false expected")
	fmt.Println(HasSuffix(s1, "iop"))
	fmt.Println(HasSuffix(s1, "dsad"))
	fmt.Println(Contains(s1, "rtyu"))
	fmt.Println(s1[9:10])
	fmt.Println(Contains(s1, "o"))
	fmt.Println(Contains(s1, "op"))
	fmt.Println(Contains(s1, "p"))
	fmt.Println(Contains(s1, "fafafas"))

}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i, _ := range s {
		if i == len(s)-len(substr)+1 {
			return false
		}
		if s[i:len(substr)+i] == substr {
			return true
		}
	}
	return false
}
