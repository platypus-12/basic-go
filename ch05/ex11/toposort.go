package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
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
	"liner algebra":        {"calculus"},
}

func main() {
	prereqs_topo_sorted, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	for i, course := range prereqs_topo_sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]string)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if seen[item] == "fin" {
				continue
			}
			if seen[item] == "tmp" {
				return errors.New("閉路があります。")
			}
			seen[item] = "tmp"
			err := visitAll(m[item])
			if err != nil {
				return err
			}
			order = append(order, item)
			seen[item] = "fin"
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	err := visitAll(keys)
	return order, err
}
