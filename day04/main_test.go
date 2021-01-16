package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 1514, Part1("test1.txt"))
	dec, sectId := getDecrypt("qzmt-zixmtkozy-ivhz-343[aaa]")
	assert.Equal(t, 343, sectId)
	assert.Equal(t, "very encrypted name", dec)
	assert.Equal(t, 993, Part2("input.txt"))
}
