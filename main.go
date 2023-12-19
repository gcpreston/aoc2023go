package main

import (
	"os"
	"slices"
	"strings"
    // "fmt"
)

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadContents(path string) string {
    dat, err := os.ReadFile(path)
    Check(err)
    return strings.TrimSpace(string(dat))
}

func Map[T, U any](arr []T, f func(T) U) []U {
    result := []U{}
    for _, t := range arr {
        result = append(result, f(t))
    }
    return result
}

func Any[T any](arr []T, f func(T) bool) bool {
    for _, t := range arr {
        if f(t) {
            return true
        }
    }
    return false
}

type Graph[T comparable] struct {
    nodes []T
    adjacencyList [][]T
}

func (g *Graph[T]) AddNode(node T) {
    g.nodes = append(g.nodes, node)
    g.adjacencyList = append(g.adjacencyList, make([]T, 0))
}

// Add an edge. Returns true if edge was added, false otherwise.
func (g *Graph[T]) AddEdge(source T, dest T) {
    i := slices.Index(g.nodes, source)
    g.adjacencyList[i] = append(g.adjacencyList[i], dest)

    j := slices.Index(g.nodes, dest)
    g.adjacencyList[j] = append(g.adjacencyList[j], source)
}

func (g Graph[T]) Neighbors(node T) []T {
    i := slices.Index(g.nodes, node)
    return g.adjacencyList[i]
}

// Traverse a graph breadth-first starting from `start`.
// Expects an operation to call for each visited node. Will continue until
// `operation` returns `true`.
func (g Graph[T]) BFSWithDepth(start T, operation func(current T, depth int) bool) {
    queue := g.Neighbors(start)
    depths := make(map[T]int)
    depths[start] = 1
    currentNode := start

    for len(queue) > 0 {
        currentNode = queue[0]
        queue = queue[1:]
        currentDepth, ok := depths[currentNode]
        if !ok {
            currentDepth = 1
        }
        // fmt.Println("bfs stop", currentNode, currentDepth)

        result := operation(currentNode, currentDepth)
        if result {
            return
        }

        for _, neighbor := range g.Neighbors(currentNode) {
            _, found := depths[neighbor]
            if !found {
                queue = append(queue, neighbor)
            }
            depths[neighbor] = currentDepth + 1
        }
    }
}

func main() {
    RunDay11()
}
