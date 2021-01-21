package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 2, Part1("test1.txt"))
}
func TestPart2(t *testing.T) {
	assert.Equal(t, 3, Part2("test2.txt"))
}
