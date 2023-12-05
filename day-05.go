package main

import (
    "fmt"
    "strings"
    "strconv"
    "slices"
)

func converter(inputVal int, destinationStart int, sourceStart int, rangeLength int) int {
    if !(inputVal >= sourceStart && inputVal <= sourceStart + rangeLength) {
        panic(fmt.Sprintf("invalid converter call: %v, %v, %v, %v\n", inputVal, destinationStart, sourceStart, rangeLength))
    }

    difference := inputVal - sourceStart
    return destinationStart + difference
}

func mapConvert(inputVal int, mappings [][]int) int {
    for _, mapping := range mappings {
        dest, source, rang := mapping[0], mapping[1], mapping[2]

        if inputVal >= source && inputVal <= source + rang {
            return converter(inputVal, dest, source, rang)
        }
    }
    return inputVal
}

func toInt(s string) int {
    dat, err := strconv.Atoi(s)
    Check(err)
    return dat
}

func RunDay05() {
    contents := ReadContents("input/day-05.txt")
    sections := strings.Split(contents, "\n\n")

    mappings := Map(sections[1:], func(section_str string) [][]int {
        lines := strings.Split(section_str, "\n")
        mapping := [][]int{}

        for _, line := range lines[1:] {
            line_data := Map(strings.Split(line, " "), toInt)
            mapping = append(mapping, line_data)
        }
        return mapping
    })

    // Part 1
    part_1_seeds := Map(strings.Split(sections[0][7:], " "), toInt)

    part_1_locations := []int{}
    for _, seed := range part_1_seeds {
        result := seed
        for _, mapping := range mappings {
            result = mapConvert(result, mapping)
        }
        part_1_locations = append(part_1_locations, result)
    }

    fmt.Println("Day 5 part 1:", slices.Min(part_1_locations))
}
