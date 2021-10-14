package main

import (
	"fmt"
)

var graph = make(map[string]map[string]bool)
var tes = make(map[string]int)

func main() {
	fmt.Println(graph)
	fmt.Println(tes)

	addEdge1("start", "end")
	fmt.Println(hasEdge("start", "end"))
	fmt.Println(hasEdge("start", "end11"))
	fmt.Println(graph)
}

func addEdge(from, to string) {
	edges := graph[from]

	if edges == nil {
		edges = make(map[string]bool)
		//graph[from] = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func addEdge1(from, to string) {
	fmt.Println("better")
	if graph[from] == nil {
		graph[from] = make(map[string]bool)
	}
	graph[from][to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
