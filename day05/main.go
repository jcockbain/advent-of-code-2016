package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	re = regexp.MustCompile(`([a-z-]+)(\d+)\[([a-z]+)\]`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("ffykfhsq"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("ffykfhsq"))
}

func Part1(doorID string) string {
	idx, res := 0, ""
	for len(res) < 8 {
		isValid, fifth, _ := isValidHash(doorID + toString(idx))
		if isValid {
			res += fifth
		}
		idx += 1
	}
	return res
}

func Part2(doorID string) string {
	idx := 0
	filled := 0
	res := make([]string, 8)
	for filled < 8 {
		isValid, sixth, seventh := isValidHash(doorID + toString(idx))
		if isValid {
			pos, err := strconv.Atoi(sixth)
			if err == nil && pos < 8 && res[pos] == "" {
				res[pos] = seventh
				filled += 1
			}
		}
		idx += 1
	}
	return strings.Join(res, "")
}

func isValidHash(inp string) (bool, string, string) {
	hash := getMD5Hash(inp)
	if startsWithFiveZeros(hash) {
		return true, string(hash[5]), string(hash[6])
	}
	return false, "a", "a"
}

func getMD5Hash(inp string) string {
	data := []byte(inp)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func startsWithFiveZeros(inp string) bool {
	if len(inp) < 5 {
		return false
	}
	for i := 0; i < 5; i++ {
		if inp[i] != '0' {
			return false
		}
	}
	return true
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

func toString(x int) string {
	return strconv.Itoa(x)
}
