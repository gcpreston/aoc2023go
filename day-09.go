package main

import (
	"strconv"
	"strings"
    "fmt"
)

func allZero(a []int) bool {
    for _, n := range a {
        if n != 0 {
            return false
        }
    }

    return true
}

func diffTree(series []int) [][]int {
    allDiffs := [][]int{series}

    for !allZero(allDiffs[len(allDiffs) - 1]) {
        currentDiffs := allDiffs[len(allDiffs) - 1]
        newDiffs := []int{}

        for i := 0; i < len(currentDiffs) - 1; i++ {
            j := i + 1
            newDiffs = append(newDiffs, currentDiffs[j] - currentDiffs[i])
        }

        allDiffs = append(allDiffs, newDiffs)
    }

    return allDiffs
}

func nextValue(tree [][]int) int {
    lastSum := 0
    for _, series := range tree {
        lastSum += series[len(series) - 1]
    }
    return lastSum
}

func nextValuePart2(tree[][]int) int {
    tree[len(tree) - 1] = append([]int{0}, tree[0]...)

    for i := len(tree) - 1; i >= 1; i-- {
        j := i - 1
        firstValInLine := tree[i][0]
        firstValInPrevLine := tree[j][0]

        tree[j] = append([]int{firstValInPrevLine - firstValInLine}, tree[j]...)
    }

    return tree[0][0]
}

func RunDay09() {
    contents := ReadContents("input/day-09.txt")
    lines := strings.Split(contents, "\n")

    data := Map(lines, func(line string) []int {
        return Map(strings.Split(line, " "), func(iStr string) int {
            n, err := strconv.Atoi(iStr)
            Check(err)
            return n
        })
    })

    total := 0
    for _, series := range data {
        tree := diffTree(series)
        next := nextValuePart2(tree)
        total += next
    }

    fmt.Println("Day 9 part 1:", total)
}
