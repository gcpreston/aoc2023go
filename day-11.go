package main

import (
    "github.com/dominikbraun/graph"
    "fmt"
    "strings"
)

type pos struct {
    x int
    y int
}

func RunDay11() {
    // Read input
    contents := ReadContents("input/day-11.txt")
    lines := strings.Split(contents, "\n")

    // Expand space
    withExpandedRows := [][]rune{}

    for _, line := range lines {
        if !strings.Contains(line, "#") {
            for i := 0; i <= 2; i++ {
                withExpandedRows = append(withExpandedRows, []rune(line))
            }
        } else {
            withExpandedRows = append(withExpandedRows, []rune(line))
        }
    }

    withExpandedCols := [][]rune{}

    for col := 0; col < len(withExpandedRows[0]); col++ {
        colChars := Map(withExpandedRows, func(row []rune) rune { return row[col] })

        if !Any(colChars, func(char rune) bool { return char == '#' }) {
            for row
        }
    }


    // Create graph

    // Find paths

    // Result

    // -------

    posHash := func(p pos) pos {
        return p
    }

    g := graph.New(posHash)

    g.AddVertex(pos{0, 0})
    g.AddVertex(pos{1, 0})
    g.AddVertex(pos{0, 1})

    g.AddEdge(pos{0, 0}, pos{1, 0})
    g.AddEdge(pos{0, 0}, pos{0, 1})
    g.AddEdge(pos{0, 0}, pos{1, 0})

    graph.DFS(g, pos{1, 0}, func(value pos) bool {
        fmt.Println(value)
        return false
    })
}
