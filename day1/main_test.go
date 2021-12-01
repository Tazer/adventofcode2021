package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountIncreaments(t *testing.T) {
	i := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	res := countIncrements(i)

	assert.Equal(t, 7, res)

}

func TestCountIncreamentsSliding(t *testing.T) {
	i := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	res := countIncrementsSliding(i)

	assert.Equal(t, 5, res)

}
