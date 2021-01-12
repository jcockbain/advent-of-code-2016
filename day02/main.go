package main

import (
	input "aoc2016/inpututils"

	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))
	fmt.Println("--- Part One ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) []int {
	lines := input.ReadLines(filename)
	c := 5
	res := []int{}

	for _, line := range lines {
		for _, i := range line {
			d := string(i)
			if d == "U" && c >= 4 {
				c -= 3
			}

			if d == "D" && c <= 6 {
				c += 3
			}

			if d == "R" && c%3 != 0 {
				c += 1
			}

			if d == "L" && c%3 != 1 {
				c -= 1
			}
		}
		res = append(res, c)
	}
	return res
}

type loc struct {
	x, y int
}

var (
	keypad = map[loc]string{
		loc{2, 0}: "1",
		loc{1, 1}: "2",
		loc{2, 1}: "3",
		loc{3, 1}: "4",
		loc{0, 2}: "5",
		loc{1, 2}: "6",
		loc{2, 2}: "7",
		loc{3, 2}: "8",
		loc{4, 2}: "9",
		loc{1, 3}: "A",
		loc{2, 3}: "B",
		loc{3, 3}: "C",
		loc{2, 4}: "D",
	}
)

func Part2(filename string) []string {
	lines := input.ReadLines(filename)
	p := loc{0, 2}
	res := []string{}

	for _, line := range lines {
		for _, i := range line {
			d := string(i)
			if d == "U" {
				ny := p.y - 1
				if _, exists := keypad[loc{p.x, ny}]; exists {
					p = loc{p.x, ny}
				}
			}

			if d == "D" {
				ny := p.y + 1
				if _, exists := keypad[loc{p.x, ny}]; exists {
					p = loc{p.x, ny}
				}
			}

			if d == "R" {
				nx := p.x + 1
				if _, exists := keypad[loc{nx, p.y}]; exists {
					p = loc{nx, p.y}
				}
			}

			if d == "L" {
				nx := p.x - 1
				if _, exists := keypad[loc{nx, p.y}]; exists {
					p = loc{nx, p.y}
				}
			}
		}
		res = append(res, keypad[p])
	}
	return res
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
