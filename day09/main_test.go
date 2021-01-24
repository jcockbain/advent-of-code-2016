package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 6, Part1("ADVENT"))
	assert.Equal(t, 7, Part1("A(1x5)BC"))
	assert.Equal(t, 9, Part1("(3x3)XYZ"))
	assert.Equal(t, 11, Part1("A(2x2)BCD(2x2)EFG"))
	assert.Equal(t, 6, Part1("(6x1)(1x3)A"))
	assert.Equal(t, 18, Part1("X(8x2)(3x3)ABCY"))
}
