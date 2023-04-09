package main

import (
	"os"
	"reflect"
	"testing"
)

// Path: algorithms/topological-sort/kahn-algorithm/kahn-algorithm.go

func Test_TopologicalSort(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)

	result, err := g.TopologicalSort()
	if err != nil {
		t.Error(err)
	}

	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
