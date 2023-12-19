package main

import (
	"fmt"
	"slices"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

type pos struct {
	row int
	col int
}

type pair struct {
	low pos
	high pos
}

func newPair(p1 pos, p2 pos) pair {
	if p1.row < p2.row || (p1.row == p2.row && p1.col < p2.col) {
		return pair{low: p1, high: p2}
	} else {
		return pair{low: p2, high: p1}
	}
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
	contents := ReadContents("input/day-11.txt")
	lines := strings.Split(contents, "\n")

	// Expand space
	expandedSpace := expandSpace(lines)

	// Create graph
	g := Graph[pos]{}
	galaxyPositions := []pos{}

	for row := 0; row < len(expandedSpace); row++ {
		for col := 0; col < len(expandedSpace[0]); col++ {
			currentPos := pos{row: row, col: col}
			g.AddNode(currentPos)

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
	foundPaths := make(map[pair]bool)
	totalCombos := len(combin.Combinations(len(galaxyPositions), 2))
	totalLen := 0

	for _, p := range galaxyPositions {
        fmt.Println("checking galaxy", p)
		if len(foundPaths) < totalCombos {
			g.BFSWithDepth(p, func(v pos, depth int) bool {
				if v != p && expandedSpace[v.row][v.col] == '#' {
					thePair := newPair(p, v)

					if !foundPaths[thePair] {
						foundPaths[thePair] = true
						totalLen += depth
					}
				}

				if len(foundPaths) == totalCombos {
                    fmt.Println("last one:", p, v)
                    return true
                }
                return false
			})
		}
	}

    fmt.Println("path count", len(foundPaths))
	// Result
	fmt.Println("Day 11 part 1:", totalLen)
}
