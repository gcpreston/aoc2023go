package main

import (
    "fmt"
    "strings"
    "strconv"
    "slices"
    "math"
    "sync"
)

func executeLine(inputVal int, destinationStart int, sourceStart int) int {
    difference := inputVal - sourceStart
    return destinationStart + difference
}

func executeMap(inputVal int, mapping [][]int) int {
    for _, line := range mapping {
        dest, source, rang := line[0], line[1], line[2]

        if inputVal >= source && inputVal <= source + rang {
            return executeLine(inputVal, dest, source)
        }
    }
    return inputVal
}

func getLocation(seed int, mappings [][][]int) int {
    result := seed
    for _, mapping := range mappings {
        result = executeMap(result, mapping)
    }
    return result
}

func getAllLocations(seeds []int, mappings [][][]int) []int {
    locations := []int{}
    for _, seed := range seeds {
        result := getLocation(seed, mappings)
        locations = append(locations, result)
    }

    return locations
}

func lowestLocation(seedRanges [][]int, mappings [][][]int) int {
    results := make(chan int)
	var wg sync.WaitGroup

    for _, seedRange := range seedRanges {
        wg.Add(1)

        go func(rang []int) {
            defer wg.Done()
            lowest := math.MaxInt
            start, length := rang[0], rang[1]

            for seed := start; seed < start + length; seed++ {
                result := getLocation(seed, mappings)
                if result < lowest {
                    lowest = result
                }
            }

            results <- lowest
        }(seedRange)
    }

    // Why does this fix the deadlock?
    go func() {
		wg.Wait()
		close(results)
	}()

    totalMin := math.MaxInt

    for testMin := range results {
        if testMin < totalMin {
            totalMin = testMin
        }
    }

    return totalMin
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

// lowest seed 3107439672? (val 52510810) (from range 2886966111 275299008)
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
    part1Seeds := Map(strings.Split(sections[0][7:], " "), toInt)
    part1Locations := getAllLocations(part1Seeds, mappings)
    fmt.Println("Day 5 part 1:", slices.Min(part1Locations))

    // Part 1.5
    fmt.Println("test thing", getLocation(3107439672, mappings))

    // Part 2
    seedRanges := chunkBy(part1Seeds, 2)
    lowest := lowestLocation(seedRanges, mappings)
    fmt.Println("Day 5 part 2:", lowest)
}
