package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := map[int][]int{}
	sLines := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		sLines = append(sLines, line)

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
}

type Line struct {
	X1, Y1, X2, Y2 int
}

type Grid struct {
	Positions map[int]map[int]int
	Lines     []Line
}

func NewGrid(lines []string) *Grid {

	dlines := []Line{}

	for _, line := range lines {
		ab := strings.Split(line, "->")

		x1, y1 := getXY(ab[0])
		x2, y2 := getXY(ab[1])

		if x1 == x2 || y1 == y2 {

			l := Line{X1: x1, Y1: y1, X2: x2, Y2: y2}

			dlines = append(dlines, l)
		}
	}

	return &Grid{
		Lines: dlines,
	}
}

func (g *Grid) MarkPositions() {
	for _, l := range g.Lines {

	}
}

func (g *Grid) GetDangerousPositions() int {
	return 0
}

func getXY(s string) (int, int) {
	s = strings.TrimSpace(s)

	x, _ := strconv.Atoi(string(s[0]))
	y, _ := strconv.Atoi(string(s[1]))

	return x, y

}
