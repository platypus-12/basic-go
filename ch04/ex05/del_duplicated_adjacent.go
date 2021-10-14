package main

import "fmt"

func del_duplicated(strings []string) []string {
	j := 0
	tmp := ""
	for i, s := range strings {
		if i == 0 || s != tmp {
			strings[j] = s
			j++
		}
		tmp = s
	}
	return strings[:j]
}

func main() {
	data := []string{"one", "one", "three"}
	data1 := []string{"1", "1", "3"}
	fmt.Printf("%q\n", del_duplicated(data[:]))
	fmt.Printf("%q\n", del_duplicated(data1[:]))
}
