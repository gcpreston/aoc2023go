package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func RunDay04() {
    dat, err := os.ReadFile("./input/day-04.txt")
    Check(err)
    contents := strings.TrimSpace(string(dat))
    lines := strings.Split(contents, "\n")
    point_total := 0
    copy_counts := map[int]int{0:1}

    for i, line := range lines {
        words := strings.Fields(strings.TrimSpace(line))
        winning_numbers := words[2:12]
        your_numbers := words[13:]
        // winning_numbers := words[2:7]
        // your_numbers := words[8:]

        win_count := 0

        for _, your_number := range your_numbers {
            if slices.Contains(winning_numbers, your_number) {
                win_count += 1
            }
        }

        if copy_counts[i] == 0 {
            copy_counts[i] = 1
        }

        for repeat := 0; repeat < copy_counts[i]; repeat++ {
            for j := i + 1; j <= i + win_count; j++ {
                if copy_counts[j] == 0 {
                    copy_counts[j] = 2
                } else {
                    copy_counts[j] += 1
                }
            }
        }

        game_points := math.Pow(2, float64(win_count - 1))
        point_total += int(game_points)

        // fmt.Println("copy counts", copy_counts)
    }

    total_cards := 0
    for _, count := range copy_counts {
        total_cards += count
    }

    fmt.Println("Day 4 part 1:", point_total)
    fmt.Println("Day 4 part 2:", total_cards)
}
