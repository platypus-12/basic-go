package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":           {"data structures"},
	"calculus":             {"liner algebra"},
	"compilers":            {"data structures", "formal languages", "computer organization"},
	"data structures":      {"discreate math"},
	"databases":            {"data structures"},
	"discreate math":       {"intro to programming"},
	"formal languages":     {"discreate math"},
	"networks":             {"operating systems"},
	"operating systems":    {"data structures", "computer organization"},
	"programming language": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string)

	visitAll = func(items map[string][]string) {
		for item, _ := range items {
			if !seen[item] {
				seen[item] = true
				item_map := make(map[string][]string)
				for _, v := range m[item] {
					item_map[v] = nil
				}
				visitAll(item_map)
				order = append(order, item)
			}
		}
	}

	visitAll(m)
	return order
}
