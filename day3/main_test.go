package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGammaAndElips(t *testing.T) {
	file, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := map[int][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		for i, l := range line {
			iL, _ := strconv.Atoi(string(l))
			if lines[i] == nil {
				lines[i] = []int{}
			}
			lines[i] = append(lines[i], iL)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res := GetGammaAndEpsilonRate(lines)

	assert.Equal(t, int64(198), res.Decimal())
}
