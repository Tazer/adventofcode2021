package main

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumbers(t *testing.T) {
	input := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"

	lines := []*Line{}

	l := parseLine(input)
	lines = append(lines, l)
	w := NewWorld(lines)
	assert.Equal(t, 1, len(w.Lines))
}

func TestBigNumbers(t *testing.T) {
	file, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []*Line{}

	for scanner.Scan() {
		line := scanner.Text()

		l := parseLine(line)
		lines = append(lines, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	w := NewWorld(lines)
	assert.Equal(t, 10, len(w.Lines))
}
