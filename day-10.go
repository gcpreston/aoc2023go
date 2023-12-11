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
    partOfLoop bool
}

func (t tile) String() string {
    return fmt.Sprintf("{row: %v, col: %v, char: %v, partOfLoop: %v}", t.row, t.col, string(t.char), t.partOfLoop)
}

func (t tile) goesLeft() bool {
    return t.partOfLoop && slices.Contains([]rune{'-', '7', 'J'}, t.char)
}

func (t tile) goesRight() bool {
    return t.partOfLoop && slices.Contains([]rune{'-', 'L', 'F'}, t.char)
}

func (t tile) goesUp() bool {
    return t.partOfLoop && slices.Contains([]rune{'|', 'L', 'J'}, t.char)
}

func (t tile) goesDown() bool {
    return t.partOfLoop && slices.Contains([]rune{'|', '7', 'F'}, t.char)
}

func canGoBetween(tileGrid [][]tile, from spaceBetween, direction string) bool {
    switch direction {
    case "up":
        if from.row == 0 {
            return false
        }

        if from.col == 0 {
            return true
        }
        if from.col == len(tileGrid[from.row]) {
            return true
        }

        tileLeft := tileGrid[from.row - 1][from.col - 1]
        tileRight := tileGrid[from.row - 1][from.col]

        return !(tileLeft.goesRight() && tileRight.goesLeft())
    case "down":
        if from.row == len(tileGrid) {
            return false
        }

        if from.col == 0 {
            return true
        }
        if from.col == len(tileGrid[from.row]) {
            return true
        }

        tileLeft := tileGrid[from.row][from.col - 1]
        tileRight := tileGrid[from.row][from.col]
        return !(tileLeft.goesRight() && tileRight.goesLeft())
    case "left":
        if from.col == 0 {
            return false
        }

        if from.row == 0 {
            return true
        }
        if from.row == len(tileGrid) {
            return true
        }
        tileUp := tileGrid[from.row - 1][from.col - 1]
        tileDown := tileGrid[from.row][from.col - 1]
        return !(tileUp.goesDown() && tileDown.goesUp())
    case "right":
        if from.col == len(tileGrid[from.row]) {
            return false
        }

        if from.row == 0 {
            return true
        }
        if from.row == len(tileGrid) {
            return true
        }
        tileUp := tileGrid[from.row - 1][from.col]
        tileDown := tileGrid[from.row][from.col]
        return !(tileUp.goesDown() && tileDown.goesUp())
    default:
        panic(fmt.Sprintf("invalid direction: %v", direction))
    }
}

// row and col here represent the space just before the row/col
// from tile.
// For example, row 1 means the space between tile row 0 and 1.
type spaceBetween struct {
    row int
    col int
}

// Find spaces that can be gone to from the current space
func (s spaceBetween) neighborSpaces(tileGrid [][]tile) []spaceBetween {
    ret := []spaceBetween{}

    if s.row > 0 {
        if canGoBetween(tileGrid, s, "up") {
            ret = append(ret, spaceBetween{row: s.row - 1, col: s.col})
        }
    }
    if s.row < len(tileGrid) - 1 {
        if canGoBetween(tileGrid, s, "down") {
            ret = append(ret, spaceBetween{row: s.row + 1, col: s.col})
        }
    }
    if s.col > 0 {
        if canGoBetween(tileGrid, s, "left") {
            ret = append(ret, spaceBetween{row: s.row, col: s.col - 1})
        }
    }
    if s.col < len(tileGrid[s.row]) - 1 {
        if canGoBetween(tileGrid, s, "right") {
            ret = append(ret, spaceBetween{row: s.row, col: s.col + 1})
        }
    }

    return ret
}

// Find outside tiles neighboring a space
func (s spaceBetween) neighborTiles(tileGrid [][]tile) []tile {
    ret := []tile{}

    spaceOnTop := s.row > 0
    spaceOnBottom := s.row < len(tileGrid)
    spaceOnLeft := s.col > 0
    spaceOnRight := s.col < len(tileGrid[s.row])

    if spaceOnTop && spaceOnLeft {
        t := tileGrid[s.row - 1][s.col - 1]
        if !t.partOfLoop {
            ret = append(ret, t)
        }
    }
    if spaceOnTop && spaceOnRight {
        t := tileGrid[s.row - 1][s.col]
        if !t.partOfLoop {
            ret = append(ret, t)
        }
    }
    if spaceOnBottom && spaceOnLeft {
        t := tileGrid[s.row][s.col - 1]
        if !t.partOfLoop {
            ret = append(ret, t)
        }
    }
    if spaceOnBottom && spaceOnRight {
        t := tileGrid[s.row][s.col]
        if !t.partOfLoop {
            ret = append(ret, t)
        }
    }

    return ret
}

func startTileConnections(tileGrid [][]tile, startTile tile) ([]tile, rune) {
    row, col := startTile.row, startTile.col
    ret := []tile{}

    connectUp := false
    connectDown := false
    connectLeft := false
    connectRight := false

    if row > 0 {
        test := tileGrid[row - 1][col]
        if slices.Contains([]rune{'|', 'F', '7'}, test.char) {
            ret = append(ret, test)
            connectUp = true
        }
    }
    if row < len(tileGrid) - 1 {
        test := tileGrid[row + 1][col]
        if slices.Contains([]rune{'|', 'L', 'J'}, test.char) {
            ret = append(ret, test)
            connectDown = true
        }
    }
    if col > 0 {
        test := tileGrid[row][col - 1]
        if slices.Contains([]rune{'-', 'L', 'F'}, test.char) {
            ret = append(ret, test)
            connectLeft = true
        }
    }
    if col < len(tileGrid[row]) - 1 {
        test := tileGrid[row][col + 1]
        if slices.Contains([]rune{'-', 'J', '7'}, test.char) {
            ret = append(ret, test)
            connectRight = true
        }
    }

    var actsAs rune
    if connectUp && connectDown {
        actsAs = '|'
    } else if connectUp && connectLeft {
        actsAs = 'J'
    } else if connectUp && connectRight {
        actsAs = 'L'
    } else if connectLeft && connectRight {
        actsAs = '-'
    } else if connectDown && connectLeft {
        actsAs = '7'
    } else if connectDown && connectRight {
        actsAs = 'F'
    }

    return ret, actsAs
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
            if char == 'S' {
                t.partOfLoop = true
                startTile = t
            }
            tileRow = append(tileRow, t)
        }
        tileGrid = append(tileGrid, tileRow)
    }

    // Iterate around loop
    conns, actsAs := startTileConnections(tileGrid, startTile)
    prevConns := []tile{startTile, startTile}
    steps := 1
    seen := []tile{startTile}

    for true {
        newConns := []tile{}
        for i, t := range conns {
            // Connection logic
            seen = append(seen, t)
            nextTile := nextConnection(tileGrid, t, prevConns[i])
            newConns = append(newConns, nextTile)

            // Mark loop
            nextTile.partOfLoop = true
            tileGrid[nextTile.row][nextTile.col] = nextTile
        }

        if Any(newConns, func (t tile) bool { return slices.Contains(seen, t) }) {
            break
        }

        prevConns = conns
        conns = newConns
        steps++
    }

    fmt.Println("Day 10 part 1:", steps)

    // Reformat start tile for part 2
    startTile.char = actsAs
    tileGrid[startTile.row][startTile.col] = startTile

    // Find outside tiles
    startSpace := spaceBetween{row: 0, col: 0}
    spacesQueue := make([]spaceBetween, 0)
    spacesQueue = append(spacesQueue, startSpace)
    visited := make(map[spaceBetween]bool)
    outsideTiles := make(map[tile]bool)

    // while the queue has spaces
    // 1. find neighboring spaces that can be moved to and not visited; add them to queue
    // 2. find neighboring tiles not part of the loop; add them to outsideTiles
    // 3. mark current space as visited
    // once complete, inside count = total - outside - loop

    for len(spacesQueue) > 0 {
        currentSpace := spacesQueue[0]
        spacesQueue = spacesQueue[1:]

        if visited[currentSpace] {
            continue
        }

        neighborSpaces := currentSpace.neighborSpaces(tileGrid)
        neighborTiles := currentSpace.neighborTiles(tileGrid)

        // fmt.Println("current space", currentSpace)
        // fmt.Println("neightbor spaces", neighborSpaces)
        // fmt.Println("neighbor tiles", neighborTiles)
        // fmt.Println()

        for _, s := range neighborSpaces {
            if !visited[s] {
                // fmt.Println("adding to queue", s)
                spacesQueue = append(spacesQueue, s)
            }
        }

        for _, t := range neighborTiles {
            outsideTiles[t] = true
        }

        visited[currentSpace] = true
    }

    totalTileCount := len(tileGrid) * len(tileGrid[0])
    outsideTileCount := len(outsideTiles)
    loopTileCount := 0

    for _, tileRow := range tileGrid {
        for _, t := range tileRow {
            if t.partOfLoop {
                loopTileCount++
            }
        }
    }

    // fmt.Printf("total %v outside %v loop %v\n", totalTileCount, outsideTileCount, loopTileCount)

    fmt.Println("Day 10 part 2:", totalTileCount - outsideTileCount - loopTileCount)
}
