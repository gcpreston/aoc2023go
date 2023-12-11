package main

import (
    "strings"
    "regexp"
    "fmt"
)

type next struct {
    l string
    r string
}

var nameRe = regexp.MustCompile(`([A-Z0-9]{3})`)

func buildNodes(lines []string) map[string]next {
    nodes := map[string]next{}

    for _, line := range lines {
        names := nameRe.FindAllString(line, -1)
        nodes[names[0]] = next{l: names[1], r: names[2]}
    }

    return nodes
}

func nextName(name string, instruction string, nodes map[string]next) string {
    next, exists := nodes[name]
    if !exists {
        panic("Node not found for name " + name)
    }

    if instruction == "L" {
        return next.l
    } else {
        return next.r
    }
}

func runPart1(instructions string, nodes map[string]next) int {
    steps := 0
    current := "AAA"

    for current != "ZZZ" {
        current = nextName(current, string(instructions[steps % len(instructions)]), nodes)
        steps++
    }

    return steps
}

func endsInZ(current string) bool {
    return current[2] == 'Z'
}

func allEndInZ(currents []string) bool {
    for _, current := range currents {
        if !endsInZ(current) {
            return false
        }
    }

    return true
}

func runPart2BruteForce(instructions string, nodes map[string]next) int {
    steps := 0
    currents := []string{}

    for name := range nodes {
        if name[2] == 'A' {
            currents = append(currents, name)
        }
    }

    for !allEndInZ(currents) {
        nextCurrents := []string{}

        for _, current := range currents {
            nextName := nextName(current, string(instructions[steps % len(instructions)]), nodes)
            nextCurrents = append(nextCurrents, nextName)
        }

        currents = nextCurrents
        steps++
    }

    return steps
}

func runPart2(instructions string, nodes map[string]next) int {
    starters := []string{}

    for name := range nodes {
        if name[2] == 'A' {
            starters = append(starters, name)
        }
    }

    loopSizes := []int{}
    for _, starter := range starters {
        loopSizes = append(loopSizes, loopSize(starter, instructions, nodes))
    }

    return multiLCM(loopSizes)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}

func multiLCM(nums []int) int {
    result := 1
    for _, num := range nums {
        result = lcm(result, num)
    }

    return result
}

func loopSize(start string, instructions string, nodes map[string]next) int {
    steps := 0
    current := start

    for !endsInZ(current) || steps == 0 {
        current = nextName(current, string(instructions[steps % len(instructions)]), nodes)
        steps++
    }

    return steps
}

func RunDay08() {
    contents := ReadContents("input/day-08.txt")
    lines := strings.Split(contents, "\n")

    instructions := lines[0]
    nodes := buildNodes(lines[2:])

    // fmt.Println("Day 8 part 1:", runPart1(instructions, nodes))
    fmt.Println("Day 8 part 2:", runPart2(instructions, nodes))
}
