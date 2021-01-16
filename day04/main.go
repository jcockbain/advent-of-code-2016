package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	input "aoc2016/inpututils"
)

var (
	re = regexp.MustCompile(`([a-z-]+)(\d+)\[([a-z]+)\]`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	totalSectorID := 0
	for _, line := range lines {
		totalSectorID += getSectorID(line)
	}
	return totalSectorID
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	for _, line := range lines {
		decrypted, sectorId := getDecrypt(line)
		if decrypted == "northpole object storage" {
			return sectorId
		}
	}
	return -1
}

func getSectorID(line string) int {
	parts := re.FindStringSubmatch(line)
	if len(parts) != 4 {
		return 0
	}
	name, sectorId, checksum := parts[1], toInt(parts[2]), parts[3]
	name = strings.ReplaceAll(name, "-", "")
	isValid := true
	mostPopularLetters := getMostPopular(name, 5)

	for _, l := range mostPopularLetters {
		if !in(checksum, l) {
			isValid = false
			break
		}
	}
	if isValid {
		return sectorId
	}
	return 0
}

func getDecrypt(line string) (string, int) {
	parts := re.FindStringSubmatch(line)
	if len(parts) != 4 {
		return "", -1
	}
	name, sectorId, _ := parts[1], toInt(parts[2]), parts[3]
	name = strings.ReplaceAll(name, "-", " ")
	decrypt := make([]rune, len(name))
	for i, c := range name {
		if c != ' ' {
			ascii := int(c) - int('a')
			newA := (ascii + sectorId) % 26
			decrypt[i] = rune(int('a') + newA)
		} else {
			decrypt[i] = ' '
		}
	}
	return strings.Trim(string(decrypt), " "), sectorId
}

type letter struct {
	char  rune
	count int
}

type lettersList []*letter

func (ll lettersList) Len() int { return len(ll) }

func (ll lettersList) Less(i, j int) bool {
	switch fd := ll[i].count - ll[j].count; {
	case fd < 0:
		return false
	case fd > 0:
		return true
	}
	return ll[i].char < ll[j].char
}

func (ll lettersList) Swap(i, j int) {
	ll[i], ll[j] = ll[j], ll[i]
}

func getMostPopular(s string, n int) []rune {
	counter := make(map[rune]int)
	for _, c := range s {
		counter[c] += 1
	}

	lfs := make(lettersList, 0)
	for l, f := range counter {
		lfs = append(lfs, &letter{l, f})
	}
	sort.Sort(lfs)

	res := make([]rune, len(lfs))
	for i, c := range lfs {
		res[i] = (*c).char
	}
	return res[:n]
}

func in(a string, b rune) bool {
	for _, c := range a {
		if c == b {
			return true
		}
	}
	return false
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
