package main

import (
    "fmt"
    "os"
    "strings"
    "regexp"
    "strconv"
)

type pull struct {
    Red, Green, Blue int
}

type game struct {
    id int
    Pulls []pull
}

const GIVEN_RED_COUNT = 12
const GIVEN_GREEN_COUNT = 13
const GIVEN_BLUE_COUNT = 14

func (g *game) IsPossible() bool {
    possible := true

    for _, p := range g.Pulls {
        if p.Red > GIVEN_RED_COUNT || p.Green > GIVEN_GREEN_COUNT || p.Blue > GIVEN_BLUE_COUNT {
            possible = false
            break
        }
    }

    return possible
}

func sum(a []int) int {
    total := 0
    for _, d := range a {
        total += d
    }
    return total
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    dat, err := os.ReadFile("./input/day-02.txt")
    check(err)
    contents := strings.TrimSpace(string(dat))

    game_re := regexp.MustCompile("Game ([0-9]+): (.+)")
    color_re := regexp.MustCompile("([0-9]+) (red|green|blue)")
    possible_ids := []int{}

    s := strings.Split(contents, "\n")
    for _, line := range s {
        game_matches := game_re.FindStringSubmatch(line)
        game_no, err := strconv.Atoi(game_matches[1])
        check(err)

        g := game{id: game_no}

        pulls_strs := strings.Split(game_matches[2], "; ")
        for _, pull_str := range pulls_strs {
            colors := strings.Split(pull_str, ", ")

            p := pull{}

            for _, color_str := range colors {
                color_matches := color_re.FindStringSubmatch(color_str)
                count, err := strconv.Atoi(color_matches[1])
                check(err)
                color := color_matches[2]

                switch color {
                case "red":
                    p.Red = count
                case "green":
                    p.Green = count
                case "blue":
                    p.Blue = count
                }
            }

            g.Pulls = append(g.Pulls, p)
        }

        if g.IsPossible() {
            possible_ids = append(possible_ids, g.id)
            fmt.Printf("possible game: %+v\n", g)
        }
    }

    total := sum(possible_ids)
    fmt.Printf("Day 2 part 1: %d\n", total)
}
