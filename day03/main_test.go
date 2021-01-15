package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, false, isValidTriangle([]int{5, 10, 15}))
	assert.Equal(t, 993, Part1("input.txt"))
	assert.Equal(t, 1849, Part2("input.txt"))
}
