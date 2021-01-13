package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	input "aoc2016/inpututils"
)

var (
	re = regexp.MustCompile(`\d+`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	valid := 0
	for _, line := range lines {
		if isValidTriangle(getSides(line)) {
			valid += 1
		}
	}
	return valid
}

func getSides(line string) []int {
	sides := re.FindAllString(line, -1)
	res := make([]int, 3)

	for i, s := range sides {
		res[i] = toInt(s)
	}

	return res
}

func isValidTriangle(sides []int) bool {
	sort.Ints(sides)
	return sides[0]+sides[1] > sides[2]
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
