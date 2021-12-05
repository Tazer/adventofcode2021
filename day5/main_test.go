package main

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	file, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g := NewGrid(lines)

	g.MarkPositions()

	log.Printf("pos: %+v", g.Positions)
	assert.Equal(t, 1, g.Positions[4][1])

	res := g.GetDangerousPositions()

	assert.Equal(t, 5, res)

}
