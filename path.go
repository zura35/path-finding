package main

import "fmt"

type node struct {
	adjacents []string
}

var (
	graph map[string]node
)

func contains(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func getAllPath(start, dest string) [][]string {
	paths := [][]string{}
	findPath(start, dest, []string{}, &paths)

	return paths
}

func findPath(start, dest string, visited []string, paths *[][]string) []string {
	visited = append(visited, start)
	if start == dest {
		*paths = append(*paths, visited)
		return visited
	}

	startNode := graph[start]
	for _, v := range startNode.adjacents {
		if contains(visited, v) {
			continue
		}
		findPath(v, dest, visited, paths)
	}

	return nil
}

func getShortestPath(start, visited []string, dest string, counter int) int {
	next := []string{}
	if len(start) == 0 {
		return -1
	}

	for _, v := range start {
		if v == dest {
			return counter
		}

		if contains(visited, v) {
			continue
		}
		next = append(next, graph[v].adjacents...)
		visited = append(visited, v)
	}

	return getShortestPath(next, visited, dest, counter+1)
}

func main() {
	a, b, c, d, e, f, g, h := node{}, node{}, node{}, node{}, node{}, node{}, node{}, node{}

	a.adjacents = []string{"B", "D", "H"}
	b.adjacents = []string{"A", "C", "D"}
	c.adjacents = []string{"B", "D", "F"}
	d.adjacents = []string{"A", "B", "C", "E"}
	f.adjacents = []string{"C", "E", "G"}
	e.adjacents = []string{"D", "F", "H"}
	g.adjacents = []string{"F", "H"}
	h.adjacents = []string{"A", "E", "G"}

	graph = make(map[string]node)
	graph["A"] = a
	graph["B"] = b
	graph["C"] = c
	graph["D"] = d
	graph["E"] = e
	graph["F"] = f
	graph["G"] = g
	graph["H"] = h

	fmt.Println(getAllPath("A", "H"))
	fmt.Println(getShortestPath([]string{"A"}, []string{}, "H", 0))
}
