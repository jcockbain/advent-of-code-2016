package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, []int{1, 9, 8, 5}, Part1("test1.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, []string{"5", "D", "B", "3"}, Part2("test1.txt"))
}
