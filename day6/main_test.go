package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFishLife26Days(t *testing.T) {
	initalState := []int{3, 4, 3, 1, 2}

	w := NewWorld(initalState, 18, false)

	res := w.Simulate()

	assert.Equal(t, 26, res)

}

func TestFishLife80(t *testing.T) {
	initalState := []int{3, 4, 3, 1, 2}
	w2 := NewWorld(initalState, 80, false)

	res2 := w2.Simulate()

	assert.Equal(t, 5934, res2)
}

func TestFishLife256(t *testing.T) {
	initalState := []int{3, 4, 3, 1, 2}
	w2 := NewWorld(initalState, 256, true)

	res2 := w2.Simulate()

	assert.Equal(t, 26984457539, res2)
}
