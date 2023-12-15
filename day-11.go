package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dominikbraun/graph"
    "gonum.org/v1/gonum/stat/combin"
)

type pos struct {
    x int
    y int
}

func (p pos) Less(comp pos) bool {
    if p.x < comp.x {
        return true
    }

    if p.y < comp.y {
        return true
    }

    return false
}

type pair struct {
    lowPos pos
    highPos pos
}

func expandSpace(space []string) []string {
    emptyRows := []int{}
    emptyCols := []int{}

    for row, line := range space {
        if !strings.Contains(line, "#") {
            emptyRows = append(emptyRows, row)
        }
    }

    for col := 0; col < len(space[0]); col++ {
        isEmpty := true

        for _, line := range space {
            if line[col] == '#' {
                isEmpty = false
            }
        }

        if isEmpty {
            emptyCols = append(emptyCols, col)
        }
    }

    var expandedSpace []string

    for row, line := range space {
        var expandedLineBuilder strings.Builder

        for col := 0; col < len(line); col++ {
            if slices.Contains(emptyCols, col) {
                expandedLineBuilder.WriteString("..")
            } else {
                expandedLineBuilder.WriteByte(line[col])
            }
        }

        expandedLine := expandedLineBuilder.String()
        expandedSpace = append(expandedSpace, expandedLine)

        if slices.Contains(emptyRows, row) {
            expandedSpace = append(expandedSpace, expandedLine)
        }
    }

    return expandedSpace
}

func RunDay11() {
    // Read input
    contents := ReadContents("input/test.txt")
    lines := strings.Split(contents, "\n")

    // Expand space
    expandedSpace := expandSpace(lines)

    // Create graph
    posHash := func(p pos) pos {
        return p
    }

    g := graph.New(posHash)
    galaxyPositions := []pos{}

    for row := 0; row < len(expandedSpace); row++ {
        for col := 0; col < len(expandedSpace[0]); col++ {
            currentPos := pos{row, col}
            g.AddVertex(currentPos)

            if expandedSpace[row][col] == '#' {
                galaxyPositions = append(galaxyPositions, currentPos)
            }

            if row > 0 {
                g.AddEdge(pos{row - 1, col}, currentPos)
            }

            if col > 0 {
                g.AddEdge(pos{row, col - 1}, currentPos)
            }
        }
    }

    // Find paths
    // inefficient
    /*
    totalLen := 0

    for _, c := range combin.Combinations(len(galaxyPositions), 2) {
        start := galaxyPositions[c[0]]
        target := galaxyPositions[c[1]]

        path, err := graph.ShortestPath(g, start, target)
        Check(err)
        totalLen += len(path) - 1
    }
    */

    // better
    foundPaths := make(map[pair]bool)
    totalCombos := len(combin.Combinations(len(galaxyPositions), 2))
    totalLen := 0

    for _, p := range galaxyPositions {
        fmt.Println("cehcking galaxy", p)
        if len(foundPaths) < totalCombos {
            graph.BFSWithDepth(g, p, func(v pos, depth int) bool {
                fmt.Println(p, v, depth)
                if v != p && expandedSpace[v.x][v.y] == '#' {
                    var thePair pair

                    if p.Less(v) {
                        thePair = pair{lowPos: p, highPos: v}
                    } else {
                        thePair = pair{lowPos: v, highPos: p}
                    }

                    if !foundPaths[thePair] {
                        fmt.Println("found a path", p, v, depth)
                        foundPaths[thePair] = true
                        totalLen += depth
                    }
                }

                return len(foundPaths) == totalCombos
            })
            return
        }
    }

    // Result
    fmt.Println("Day 11 part 1:", totalLen)
}
