package main

import (
	"fmt"

	input "aoc2016/inpututils"
	"regexp"
	"strconv"
)

var (
	re1 = regexp.MustCompile(`rect (\d+)x(\d+)`)
	re2 = regexp.MustCompile(`rotate (row|column) [x|y]=(\d+) by (\d+)`)
)

type pos struct {
	x, y int
}

type points map[pos]int

type screen struct {
	h, w int
	p    points
}

func (s screen) countOn() int {
	res := 0
	for _, on := range s.p {
		if on == 1 {
			res += 1
		}
	}
	return res
}

func (s screen) printScreen() {
	points := s.p
	for y := 0; y < s.h; y++ {
		row := ""
		for x := 0; x < s.w; x++ {
			if points[pos{x, y}] == 1 {
				row += "#"
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
	fmt.Println("<-------->")
}

func newScreen(h int, w int) screen {
	p := points{}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p[pos{x, y}] = 0
		}
	}
	return screen{
		h,
		w,
		p,
	}
}

func main() {
	res := Part1("input.txt", 6, 50)
	fmt.Println("--- Part One ---")
	fmt.Println(res.countOn())
	fmt.Println("--- Part Two ---")
	res.printScreen()
}

func Part1(filename string, h int, w int) screen {
	inp := input.ReadLines(filename)
	s := newScreen(h, w)
	for _, line := range inp {
		if re1.MatchString(line) {
			parts := re1.FindStringSubmatch(line)
			x, y := toInt(parts[1]), toInt(parts[2])
			newPoints := points{}
			for p, on := range s.p {
				if p.x < x && p.y < y {
					newPoints[p] = 1
				} else {
					newPoints[p] = on
				}
			}
			s.p = newPoints
		}
		if re2.MatchString(line) {
			parts := re2.FindStringSubmatch(line)
			dim, i, v := parts[1], toInt(parts[2]), toInt(parts[3])
			newPoints := points{}
			for p, on := range s.p {
				if dim == "row" && p.y == i {
					newPoints[pos{(p.x + v) % s.w, p.y}] = on
				} else if dim == "column" && p.x == i {
					newPoints[pos{p.x, (p.y + v) % s.h}] = on
				} else {
					newPoints[p] = on
				}
			}
			s.p = newPoints
		}
	}
	return s
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
