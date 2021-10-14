package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	//out := strings[:0]
	out := []string{}
	//var out []string
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	data1 := []string{"1", "", "3"}
	fmt.Printf("%q\n", nonempty(data[:]))
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", nonempty2(data1[:]))
	fmt.Printf("%q\n", data1)
}
