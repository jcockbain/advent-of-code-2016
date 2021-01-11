package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 5, Part1("test1.txt"))
	assert.Equal(t, 2, Part1("test2.txt"))
	assert.Equal(t, 12, Part1("test3.txt"))
}
func TestPart2(t *testing.T) {
	assert.Equal(t, 4, Part2("test4.txt"))
}
