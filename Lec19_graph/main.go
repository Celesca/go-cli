package main

import "fmt"

type Graph struct {
	numVertices int
	adjList     [][]int
}

func NewGraph(n int) *Graph {
	adjList := make([][]int, n)
	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0)
	}
	return &Graph{numVertices: n, adjList: adjList}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) PrintGraph() {
	for i := 0; i < g.numVertices; i++ {
		fmt.Printf("%d: ", i)
		for j := 0; j < len(g.adjList[i]); j++ {
			fmt.Printf("%d ", g.adjList[i][j])
		}
		fmt.Println()
	}
}

func main() {
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.PrintGraph()

}
