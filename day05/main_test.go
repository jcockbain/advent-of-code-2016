package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	isValid, letter, _ := isValidHash("abc3231929")
	assert.Equal(t, true, isValid)
	assert.Equal(t, "1", letter)
	assert.Equal(t, "18f47a30", Part1("abc"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "05ace8e3", Part2("abc"))
}
