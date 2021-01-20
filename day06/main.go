package main

import (
	"fmt"

	input "aoc2016/inpututils"
	"math"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) string {
	lines := input.ReadLines(filename)
	lineLen := len(lines[0])
	counter := make(map[int]map[rune]int, 0)

	for i := 0; i < lineLen; i++ {
		counter[i] = map[rune]int{}
	}

	for _, line := range lines {
		for i, c := range line {
			counter[i][c] += 1
		}
	}

	res := ""
	for i := 0; i < lineLen; i++ {
		res += string(getMaxKey(counter[i]))
	}
	return res
}

func Part2(filename string) string {
	lines := input.ReadLines(filename)
	lineLen := len(lines[0])
	counter := make(map[int]map[rune]int, 0)

	for i := 0; i < lineLen; i++ {
		counter[i] = map[rune]int{}
	}

	for _, line := range lines {
		for i, c := range line {
			counter[i][c] += 1
		}
	}

	res := ""
	for i := 0; i < lineLen; i++ {
		res += string(getMinKey(counter[i]))
	}
	return res
}

func getMaxKey(m map[rune]int) rune {
	maxVal, maxKey := 0, 'a'
	for k, v := range m {
		if v > maxVal {
			maxVal = v
			maxKey = k
		}
	}
	return maxKey
}

func getMinKey(m map[rune]int) rune {
	minVal, minKey := math.MaxInt32, 'a'
	for k, v := range m {
		if v < minVal {
			minVal = v
			minKey = k
		}
	}
	return minKey
}
