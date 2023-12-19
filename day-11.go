package main

import (
	"fmt"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

const INPUT_PATH = "input/day-11.txt"
const EXPAND_SIZE = 1000000

type pos struct {
	row int
	col int
}

func findExpandedSpace(space []string) ([]int, []int) {
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

	return emptyRows, emptyCols
}

func expandedPathLength(p1 pos, p2 pos, expandedRows []int, expandedCols []int, expandSize int) int {
    rowCount := 0
    colCount := 0

    lowRow := min(p1.row, p2.row)
    highRow := max(p1.row, p2.row)
    lowCol := min(p1.col, p2.col)
    highCol := max(p1.col, p2.col)

    for _, r := range expandedRows {
        if r > lowRow && r < highRow {
            rowCount++
        }
    }

    for _, c := range expandedCols {
        if c > lowCol && c < highCol {
            colCount++
        }
    }

    return (rowCount + colCount) * (expandSize - 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func RunDay11() {
	// Read input
	contents := ReadContents(INPUT_PATH)
	lines := strings.Split(contents, "\n")

	// Expand space
	expandedRows, expandedCols := findExpandedSpace(lines)

	// Create graph
	galaxyPositions := []pos{}

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[0]); col++ {
			currentPos := pos{row: row, col: col}

			if lines[row][col] == '#' {
				galaxyPositions = append(galaxyPositions, currentPos)
			}
		}
	}

	// Find paths
    baseLen := 0
    expandLen := 0

    fmt.Println("rows", expandedRows)
    fmt.Println("cols", expandedCols)

    for _, c := range combin.Combinations(len(galaxyPositions), 2) {
        gal1 := galaxyPositions[c[0]]
        gal2 := galaxyPositions[c[1]]

        shortestDistance := abs(gal1.row - gal2.row) + abs(gal1.col - gal2.col)
        baseLen += shortestDistance
        expandedDistance := expandedPathLength(gal1, gal2, expandedRows, expandedCols, EXPAND_SIZE)
        expandLen += expandedDistance
    }

	// Result
	fmt.Println("Day 11 part 2:", baseLen + expandLen)
}
