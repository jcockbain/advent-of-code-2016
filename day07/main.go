package main

import (
	"fmt"

	input "aoc2016/inpututils"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) int {
	inp := input.ReadLines(filename)
	res := 0
	for _, line := range inp {
		if supportsTLS(line) {
			res += 1
		}
	}
	return res
}

func supportsTLS(line string) bool {
	inBrackets := false
	seenOutsideBrackets := false
	for i := 0; i < len(line)-3; i++ {
		if line[i] == '[' {
			inBrackets = true
			continue
		}

		if line[i] == ']' {
			inBrackets = false
			continue
		}

		if line[i] == line[i+3] && line[i+1] == line[i+2] && line[i] != line[i+1] {
			if inBrackets {
				return false
			} else {
				seenOutsideBrackets = true
			}
		}
	}
	return seenOutsideBrackets
}

func Part2(filename string) int {
	inp := input.ReadLines(filename)
	res := 0
	for _, line := range inp {
		if supportsSSL(line) {
			res += 1
		}
	}
	return res
}

func supportsSSL(line string) bool {
	inBrackets := false
	babSeen := map[string]bool{}
	abaSeen := map[string]bool{}

	for i := 0; i < len(line)-2; i++ {
		if line[i] == '[' {
			inBrackets = true
			continue
		}
		if line[i] == ']' {
			inBrackets = false
			continue
		}

		if line[i] == line[i+2] && line[i] != line[i+1] {

			if inBrackets {
				babSeen[string([]byte{line[i], line[i+1]})] = true
			} else {
				abaSeen[string([]byte{line[i], line[i+1]})] = true
			}
		}
	}

	for s1, _ := range abaSeen {
		for s2, _ := range babSeen {
			if s1[0] == s2[1] && s1[1] == s2[0] {
				return true
			}
		}
	}

	return false
}
