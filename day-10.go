package main

import (
    "strings"
    "slices"
    "fmt"
)

type tile struct {
    row int
    col int
    char rune
}

func (t tile) String() string {
    return fmt.Sprintf("{row: %v, col: %v, char: %v}", t.row, t.col, string(t.char))
}

type loopPiece struct {
    conn1 *tile
    conn2 *tile
}

type loop struct {
    start loopPiece
}

func startTileConnections(tileGrid [][]tile, startTile tile) []tile {
    row, col := startTile.row, startTile.col
    ret := []tile{}

    if row > 0 {
        test := tileGrid[row - 1][col]
        if slices.Contains([]rune{'|', 'F', '7'}, test.char) {
            ret = append(ret, test)
        }
    }
    if row < len(tileGrid) - 1 {
        test := tileGrid[row + 1][col]
        if slices.Contains([]rune{'|', 'L', 'J'}, test.char) {
            ret = append(ret, test)
        }
    }
    if col > 0 {
        test := tileGrid[row][col - 1]
        if slices.Contains([]rune{'-', 'L', 'F'}, test.char) {
            ret = append(ret, test)
        }
    }
    if col < len(tileGrid[row]) - 1 {
        test := tileGrid[row][col + 1]
        if slices.Contains([]rune{'-', 'J', '7'}, test.char) {
            ret = append(ret, test)
        }
    }

    return ret
}

// Should return an array of size 2 for any tile within the loop
func loopConnections(tileGrid [][]tile, current tile) []tile {
    switch current.char {
    case '|':
        return []tile{tileGrid[current.row - 1][current.col], tileGrid[current.row + 1][current.col]}
    case '-':
        return []tile{tileGrid[current.row][current.col - 1], tileGrid[current.row][current.col + 1]}
    case 'L':
        return []tile{tileGrid[current.row - 1][current.col], tileGrid[current.row][current.col + 1]}
    case 'J':
        return []tile{tileGrid[current.row - 1][current.col], tileGrid[current.row][current.col - 1]}
    case 'F':
        return []tile{tileGrid[current.row + 1][current.col], tileGrid[current.row][current.col + 1]}
    case '7':
        return []tile{tileGrid[current.row + 1][current.col], tileGrid[current.row][current.col - 1]}
    default:
        panic(fmt.Sprintf("Not part of the loop: %+v", current))
    }
}

func nextConnection(tileGrid [][]tile, current tile, prev tile) tile {
    conns := loopConnections(tileGrid, current)
    for _, t := range conns {
        if t.row != prev.row || t.col != prev.col {
            return t
        }
    }
    // fmt.Printf("tile: %v, conns: %v, prev: %v\n", current, conns, prev)
    panic(fmt.Sprintf("Couldn't find next connections of tile: %+v", current))
}

func RunDay10() {
    contents := ReadContents("input/day-10.txt")
    lines := strings.Split(contents, "\n")

    // Build data
    tileGrid := [][]tile{}
    var startTile tile

    for row, line := range lines {
        tileRow := []tile{}
        for col, char := range line {
            t := tile{row: row, col: col, char: char}
            tileRow = append(tileRow, t)
            if char == 'S' {
                startTile = t
            }
        }
        tileGrid = append(tileGrid, tileRow)
    }

    // Iterate around loop
    conns := startTileConnections(tileGrid, startTile)
    prevConns := []tile{startTile, startTile}
    steps := 1
    seen := []tile{startTile}

    for true {
        newConns := []tile{}
        for i, t := range conns {
            seen = append(seen, t)
            newConns = append(newConns, nextConnection(tileGrid, t, prevConns[i]))
        }

        if Any(newConns, func (t tile) bool { return slices.Contains(seen, t) }) {
            break
        }

        prevConns = conns
        conns = newConns
        steps++
    }

    fmt.Println("Day 10 part 1:", steps)
}
