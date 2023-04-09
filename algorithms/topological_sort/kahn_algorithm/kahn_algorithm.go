package main

import (
	"errors"
	"fmt"
)

type Graph struct {
	edges    map[int][]int
	indegree map[int]int
}

func NewGraph() *Graph {
	return &Graph{
		edges:    make(map[int][]int),
		indegree: make(map[int]int),
	}
}

/*
Time Complexity: O(V+E).
	The outer for loop will be executed V number of times and the inner for loop will be executed E number of times.
Auxiliary Space: O(V).
	The queue needs to store all the vertices of the graph. So the space required is O(V)
*/

func (g *Graph) AddEdge(from, to int) {
	g.edges[from] = append(g.edges[from], to)

	// go map: if key not exist, init its value to 0
	if g.indegree[from] == 0 {
		g.indegree[from] = 0
	}

	g.indegree[to]++
}

func (g *Graph) TopologicalSort() ([]int, error) {
	result := make([]int, 0)
	queue := make([]int, 0)

	// find all nodes with in-degree 0 and add them to the queue
	for k, v := range g.indegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		// using slice to simulate queue
		u := queue[0]
		queue = queue[1:]

		result = append(result, u)
		for _, v := range g.edges[u] {
			g.indegree[v]--
			if g.indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(result) != len(g.indegree) {
		return nil, errors.New("graph contains cycle")
	}

	return result, nil
}

func main() {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	result, err := g.TopologicalSort()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
