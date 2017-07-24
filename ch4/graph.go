/*
 * Compilation: go build graph.go
 * Execution: ./graph
 *
 * The value typ e of a map can its elf be a composite typ e,
 * such as a map or slice.
 *
 * $ ./graph
 * show graph:  map[topleft:map[red:true] topright:map[yellow:true]
 * buttomleft:map[green:true] buttomright:map[blue:true]]
 * hasEdge topleft:red true
 * hasEdge topleft:yellow false
 */

package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("topleft", "red")
	addEdge("topright", "yellow")
	addEdge("buttomleft", "green")
	addEdge("buttomright", "blue")
	fmt.Println("show graph: ", graph)
	fmt.Println("hasEdge topleft:red", hasEdge("topleft", "red"))
	fmt.Println("hasEdge topleft:yellow", hasEdge("topleft", "yellow"))
}
