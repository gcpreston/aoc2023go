package main

import (
    "fmt"
    "os"
    "strings"
    "regexp"
    "strconv"
    "unicode"
)

func getSubstring(s string, indices []int) string {
    return string(s[indices[0]:indices[1]])
}

func getSurroundings(input []string, row int, indices []int) string {
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

    return surroundings.String()
}

var symbolRe = regexp.MustCompile(`[^\.0-9]`)

func isEnginePart(input []string, row int, indices []int) bool {
    surroundings := getSurroundings(input, row, indices)
    return symbolRe.MatchString(surroundings)
}

func numberFromPosition(line string, col int) int {
    start, end := col, col

    for start - 1 >= 0 && unicode.IsDigit(rune(line[start - 1])) {
        start -= 1
    }

    for end < len(line) && unicode.IsDigit(rune(line[end])) {
        end += 1
    }

    // fmt.Printf("numberFromPosition line %v col %v: start %v end %v\n", line, col, start, end)

    num, err := strconv.Atoi(line[start:end])
    Check(err)
    return num
}

var numberRe = regexp.MustCompile("([0-9]+)")

func gearRatio(input []string, row int, indices []int) int {
    // 1. Find the numbers surrounding the gear
    // 2. Multiply them together

    colStart := max(indices[0] - 1, 0)
    colEnd := min(indices[1] + 1, len(input[row]) - 1)

    nums := []int{}

    if row > 0 {
        chunk := getSubstring(input[row - 1], []int{colStart, colEnd})
        positions := numberRe.FindAllStringIndex(chunk, -1)

        if positions != nil {
            for _, position := range positions {
                nums = append(nums, numberFromPosition(input[row - 1], colStart + position[0]))
            }
        }
    }

    if indices[0] > 0 {
        if unicode.IsDigit(rune(input[row][indices[0] - 1])) {
            nums = append(nums, numberFromPosition(input[row], indices[0] - 1))
        }
    }

    if indices[1] < (len(input[row]) - 1) {
        if unicode.IsDigit(rune(input[row][indices[1]])) {
            nums = append(nums, numberFromPosition(input[row], indices[1]))
        }
    }

    if row < (len(input) - 1) {
        chunk := getSubstring(input[row + 1], []int{colStart, colEnd})
        positions := numberRe.FindAllStringIndex(chunk, -1)

        if positions != nil {
            for _, position := range positions {
                nums = append(nums, numberFromPosition(input[row + 1], colStart + position[0]))
            }
        }
    }

    fmt.Printf("gear row %v col %v nums %v\n", row, indices[0], nums)

    if len(nums) == 2 {
        return nums[0] * nums[1]
    }

    return 0
}

func RunDay03() {
    dat, err := os.ReadFile("./input/day-03.txt")
    Check(err)
    contents := strings.TrimSpace(string(dat))
    lines := strings.Split(contents, "\n")

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

    // Part 1
    total := 0

    for row, line := range lines {
        result := numberRe.FindAllStringIndex(line, -1)

        for _, matchIndices := range result {
            numStr := getSubstring(line, matchIndices)

            // fmt.Printf("is %v an engine part? %v\n", numStr, isEnginePart(lines, row, matchIndices))

            if isEnginePart(lines, row, matchIndices) {
                num, err := strconv.Atoi(numStr)
                Check(err)
                total += num
            }
        }
    }

    fmt.Printf("Day 3 part 1: %v\n", total)

    // Part 2
    gearRe := regexp.MustCompile(`\*`)
    gearRatioTotal := 0

    for row, line := range lines {
        result := gearRe.FindAllStringIndex(line, -1)

        for _, matchIndices := range result {
            ratio := gearRatio(lines, row, matchIndices)
            gearRatioTotal += ratio
        }
    }

    fmt.Printf("Day 3 part 2: %v\n", gearRatioTotal)
}
