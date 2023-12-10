package main

import (
    "os"
    "strings"
)

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadContents(path string) string {
    dat, err := os.ReadFile(path)
    Check(err)
    return strings.TrimSpace(string(dat))
}

func Map[T, U any](arr []T, f func(T) U) []U {
    result := []U{}
    for _, t := range arr {
        result = append(result, f(t))
    }
    return result
}

func Any[T any](arr []T, f func(T) bool) bool {
    for _, t := range arr {
        if f(t) {
            return true
        }
    }
    return false
}

func main() {
    RunDay10()
}
