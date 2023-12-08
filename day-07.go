package main

import (
    "strings"
    "strconv"
    "sort"
    "fmt"
)

var cardValuesMap = map[rune]int{
    'J': 1,
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
    'T': 10,
    // 'J': 11,
    'Q': 12,
    'K': 13,
    'A': 14,
}

const CARDS_PER_HAND = 5

// Type values:
// High card: 1
// One pair: 2
// Two pair: 3
// Three of a kind: 4
// Full house: 5
// Four of a kind: 6
// Five of a kind: 7

type hand struct {
    typeValue int
    cardValues []int
}

type handAndBid struct {
    hand hand
    bid int
}

func testEq(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func handTypeFromBuckets(buckets map[int]int) int {
    counts := make([]int, 0, len(buckets))
    for _, val := range buckets {
        counts = append(counts, val)
    }

    sort.Ints(counts)

    if testEq(counts, []int{5}) {
        return 7
    } else if testEq(counts, []int{1, 4}) {
        return 6
    } else if testEq(counts, []int{2, 3}) {
        return 5
    } else if testEq(counts, []int{1, 1, 3}) {
        return 4
    } else if testEq(counts, []int{1, 2, 2}) {
        return 3
    } else if testEq(counts, []int{1, 1, 1, 2}) {
        return 2
    } else {
        return 1
    }
}

// Transform Js to whatever will produce the highest value type
func transformBuckets(buckets map[int]int) map[int]int {
    if jCount, ok := buckets[1]; ok {
        delete(buckets, 1)

        cardValToChangeTo, maxCount := 0, 0
        for cardValue, count := range buckets {
            if count > maxCount {
                cardValToChangeTo = cardValue
                maxCount = count
            }
        }

        buckets[cardValToChangeTo] += jCount
    }

    return buckets
}

func newHand(cardsStr string) hand {
    cardValues := []int{}
    buckets := make(map[int]int)

    // Calculate card values and pairings
    for _, cardRune := range cardsStr {
        val := cardValuesMap[cardRune]
        cardValues = append(cardValues, val)
        buckets[val] += 1
    }

    // Calculate hand type based on pairings
    buckets = transformBuckets(buckets) // Comment this out for part 1
    typeValue := handTypeFromBuckets(buckets)

    return hand{cardValues: cardValues, typeValue: typeValue}
}

func (h1 hand) isStrongerThan(h2 hand) bool {
    if h1.typeValue > h2.typeValue {
        return true
    } else if h1.typeValue < h2.typeValue {
        return false
    }

    for i := 0; i < CARDS_PER_HAND; i++ {
        if h1.cardValues[i] > h2.cardValues[i] {
            return true
        } else if h1.cardValues[i] < h2.cardValues[i] {
            return false
        }
    }

    return false
}

func (h hand) toString() string {
    var typeRepr string

    switch h.typeValue {
    case 1:
        typeRepr = "High card"
    case 2:
        typeRepr = "One pair"
    case 3:
        typeRepr = "Two pair"
    case 4:
        typeRepr = "Three of a kind"
    case 5:
        typeRepr = "Full house"
    case 6:
        typeRepr = "Four of a kind"
    case 7:
        typeRepr = "Five of a kind"
    }

    return fmt.Sprintf("{type: %v, cardValues: %v}", typeRepr, h.cardValues)
}

func RunDay07() {
    // contents := ReadContents("input/test.txt")
    contents := ReadContents("input/day-07.txt")
    lines := strings.Split(contents, "\n")
    handsAndBids := []handAndBid{}

    for _, line := range lines {
        splt := strings.Split(strings.TrimSpace(line), " ")
        cardsStr, bidStr := splt[0], splt[1]

        hand := newHand(cardsStr)
        bid, err := strconv.Atoi(bidStr)
        Check(err)

        handsAndBids = append(handsAndBids, handAndBid{hand: hand, bid: bid})
        // fmt.Printf("L: %v, P: %v\n", line, hand.toString())
    }

    sort.Slice(handsAndBids, func(i, j int) bool {
		return handsAndBids[j].hand.isStrongerThan(handsAndBids[i].hand)
	})

    totalWinnings := 0
    for i, hnb := range handsAndBids {
        totalWinnings += hnb.bid * (i + 1)
    }

    fmt.Println("Day 7 part 2:", totalWinnings)
}
