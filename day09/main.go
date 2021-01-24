package main

import (
	"fmt"

	input "aoc2016/inpututils"
	"regexp"
	"strconv"
)

var (
	re1 = regexp.MustCompile(`(\d+)x(\d+)`)
)

func main() {
	inp := input.ReadRaw("input.txt")
	fmt.Println("--- Part One ---")
	fmt.Println(Part1(inp))
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func Part1(inp string) int {
	res, i := "", 0
	for i < len(inp) {
		c := inp[i]
		if c == '(' {
			marker := ""
			i += 1
			for inp[i] != ')' {
				marker += string(inp[i])
				i += 1
			}
			parts := re1.FindStringSubmatch(marker)
			repeatLen, repeatTimes := toInt(parts[1]), toInt(parts[2])
			i += 1
			repeatStr := inp[i : i+repeatLen]
			for t := 0; t < repeatTimes; t++ {
				res += repeatStr
			}
			i += repeatLen
		} else {
			res += string(c)
			i += 1
		}
	}
	return len(res)
}
