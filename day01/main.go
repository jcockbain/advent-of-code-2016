package main

import (
	input "aoc2016/inpututils"

	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	re1 = regexp.MustCompile(`([A-Z])(\d+)`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) int {
	dirs := input.ReadRaw(filename)
	s := newStreets()
	for _, c := range strings.Split(dirs, ", ") {
		parts := re1.FindStringSubmatch(c)
		d, steps := parts[1], toInt(parts[2])

		if d == "L" {
			if s.dir == 0 {
				s.dir = 3
			} else {
				s.dir = s.dir - 1
			}
		} else if d == "R" {
			s.dir = (s.dir + 1) % 4
		}

		for i := 0; i < steps; i++ {
			s.move()
		}

	}

	return Abs(s.pos.x) + Abs(s.pos.y)
}

type loc struct {
	x, y int
}

type streets struct {
	visited map[loc]bool
	pos     loc
	dir     int
}

func newStreets() streets {
	return streets{
		map[loc]bool{},
		loc{0, 0},
		0,
	}
}

func (s *streets) move() {
	if s.dir == 0 {
		s.pos.y += 1
	}
	if s.dir == 1 {
		s.pos.x += 1
	}
	if s.dir == 2 {
		s.pos.y -= 1
	}
	if s.dir == 3 {
		s.pos.x -= 1
	}
}

func Part2(filename string) int {
	dirs := input.ReadRaw(filename)
	s := newStreets()
	instructions := strings.Split(dirs, ", ")
	for _, c := range instructions {
		parts := re1.FindStringSubmatch(c)
		d, steps := parts[1], toInt(parts[2])

		if d == "L" {
			if s.dir == 0 {
				s.dir = 3
			} else {
				s.dir = s.dir - 1
			}
		} else if d == "R" {
			s.dir = (s.dir + 1) % 4
		}

		for i := 0; i < steps; i++ {
			s.visited[s.pos] = true
			s.move()
			if isIn(s.visited, s.pos) {
				return Abs(s.pos.x) + Abs(s.pos.y)
			}
		}
	}
	return -1
}

func isIn(m map[loc]bool, x loc) bool {
	_, t := m[x]
	return t
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
