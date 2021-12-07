package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuel(t *testing.T) {
	initialState := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	w := NewWorld(initialState)

	res := w.Simulate(false)

	assert.Equal(t, 37, res)

	res = w.Simulate(true)

	assert.Equal(t, 168, res)
}
