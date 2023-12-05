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

func getAllLocations(seeds []int, mappings [][][]int) []int {
    locations := []int{}
    for _, seed := range seeds {
        result := seed
        for _, mapping := range mappings {
            result = mapConvert(result, mapping)
        }
        locations = append(locations, result)
    }

    return locations
}

func toInt(s string) int {
    dat, err := strconv.Atoi(s)
    Check(err)
    return dat
}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
    for chunkSize < len(items) {
        items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
    }
    return append(chunks, items)
}

func RunDay05() {
    contents := ReadContents("input/test.txt")
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
    part1Seeds := Map(strings.Split(sections[0][7:], " "), toInt)
    part1Locations := getAllLocations(part1Seeds, mappings)

    fmt.Println("Day 5 part 1:", slices.Min(part1Locations))

    // Part 2
    seedRanges := chunkBy(part1Seeds, 2)
    part2Seeds := []int{}

    for _, seedRange := range seedRanges {
        start, length := seedRange[0], seedRange[1]

        for seed := start; seed <= start + length; seed++ {
            part2Seeds = append(part2Seeds, seed)
        }
    }

    part2Locations := getAllLocations(part2Seeds, mappings)
    fmt.Println("Day 5 part 2:", slices.Min(part2Locations))
}
