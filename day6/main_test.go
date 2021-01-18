package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, "easter", Part1("test.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "advent", Part2("test.txt"))
}
