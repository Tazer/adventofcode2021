package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	input := []Movement{
		{direction: "forward", steps: 5},
		{direction: "down", steps: 5},
		{direction: "forward", steps: 8},
		{direction: "up", steps: 3},
		{direction: "down", steps: 8},
		{direction: "forward", steps: 2},
	}

	res := CalculateMovements(input, false)

	assert.Equal(t, 15, res.horizontal)
	assert.Equal(t, 10, res.depth)
	assert.Equal(t, 150, res.Result())

	res2 := CalculateMovements(input, true)

	assert.Equal(t, 15, res2.horizontal)
	assert.Equal(t, 60, res2.depth)
	assert.Equal(t, 900, res2.Result())
}
