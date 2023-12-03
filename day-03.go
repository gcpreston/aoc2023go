package main

import (
    "fmt"
    "os"
    "strings"
    "regexp"
    "strconv"
)

func getSubstring(s string, indices []int) string {
    return string(s[indices[0]:indices[1]])
}

var symbolRe = regexp.MustCompile(`[^\.0-9]`)

func isEnginePart(input []string, row int, indices []int) bool {
    var surroundings strings.Builder

    colStart := max(indices[0] - 1, 0)
    colEnd := min(indices[1] + 1, len(input[row]) - 1)

    if row > 0 {
        surroundings.WriteString(getSubstring(input[row - 1], []int{colStart, colEnd}))
    }

    if indices[0] > 0 {
        surroundings.WriteByte(input[row][indices[0] - 1])
    }

    if indices[1] < (len(input[row]) - 1) {
        surroundings.WriteByte(input[row][indices[1]])
    }

    if row < (len(input) - 1) {
        surroundings.WriteString(getSubstring(input[row + 1], []int{colStart, colEnd}))
    }

    if symbolRe.MatchString(surroundings.String()) {
        return true
    }

    return false
}

func RunDay03() {
    dat, err := os.ReadFile("./input/day-03.txt")
    Check(err)
    contents := strings.TrimSpace(string(dat))
    lines := strings.Split(contents, "\n")

    numberRe := regexp.MustCompile("([0-9]+)")

    // 1. Iterate over lines
    // 2. Find index of start of number and length
    // 3. Check surrounding indices for symbol (including next lines)
    // 4. Repeat until line is finished
    // 5. Repeat until out of lines


    // indices are column
    // row - 1, colStart - 1 => colEnd + 1
    // row, colStart - 1
    // row, colEnd + 1
    // row + 1, colStart - 1 => colEnd + 1

    total := 0

    for row, line := range lines {
        result := numberRe.FindAllStringIndex(line, -1)

        for _, matchIndices := range result {
            numStr := getSubstring(line, matchIndices)

            fmt.Printf("is %v an engine part? %v\n", numStr, isEnginePart(lines, row, matchIndices))

            if isEnginePart(lines, row, matchIndices) {
                num, err := strconv.Atoi(numStr)
                Check(err)
                total += num
            }
        }
    }

    fmt.Printf("Day 3 part 1: %v\n", total)
}
