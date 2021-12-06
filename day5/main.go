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

	g := NewGrid(sLines)
	g.MarkPositions(false)

	res := g.GetDangerousPositions()

	log.Printf("res %d", res)

	g = NewGrid(sLines)
	g.MarkPositions(true)

	res = g.GetDangerousPositions()

	log.Printf("res2 %d", res)
}

type Line struct {
	X1, Y1, X2, Y2 int
}

type Grid struct {
	Positions map[int]map[int]int
	Lines     []Line
	DLines    []Line
}

func NewGrid(lines []string) *Grid {

	slines := []Line{}
	dlines := []Line{}

	for _, line := range lines {
		ab := strings.Split(line, "->")
		x1, y1 := getXY(ab[0])
		x2, y2 := getXY(ab[1])

		if x1 == x2 || y1 == y2 {

			l := Line{X1: x1, Y1: y1, X2: x2, Y2: y2}
			slines = append(slines, l)
		} else if x1 == y1 && x2 == y2 || x1 == y2 && x2 == y1 || y2%x1 == 0 && y1%x2 == 0 {
			l := Line{X1: x1, Y1: y1, X2: x2, Y2: y2}
			log.Printf("adding diagonal line %v", l)
			dlines = append(dlines, l)
		} else {
			log.Printf("skipping line [%d,%d -> %d,%d]", x1, y1, x2, y2)
		}
	}

	return &Grid{
		Lines:  slines,
		DLines: dlines,
	}
}

func (g *Grid) MarkPositions(dLines bool) {
	for _, l := range g.Lines {
		if l.X1 == l.X2 {

			big := l.Y2
			small := l.Y1

			if l.Y1 > l.Y2 {
				big = l.Y1
				small = l.Y2
			}

			for y := small; y <= big; y++ {
				if g.Positions == nil {
					g.Positions = map[int]map[int]int{}
				}
				if g.Positions[y] == nil {
					g.Positions[y] = map[int]int{}
				}
				g.Positions[y][l.X1] += 1
			}
		} else {

			big := l.X2
			small := l.X1

			if l.X1 > l.X2 {
				big = l.X1
				small = l.X2
			}

			for x := small; x <= big; x++ {
				if g.Positions == nil {
					g.Positions = map[int]map[int]int{}
				}
				if g.Positions[l.Y1] == nil {
					g.Positions[l.Y1] = map[int]int{}
				}
				g.Positions[l.Y1][x] += 1
			}
		}
	}
	if dLines {
		log.Printf("whats dlines %+v", g.DLines)
		for _, l := range g.DLines {
			ybig := l.Y2
			ysmall := l.Y1
			startX := l.X1

			if l.Y1 > l.Y2 {
				ybig = l.Y1
				ysmall = l.Y2
				startX = l.X2
			}

			//xbig := l.X2
			// xsmall := l.X1

			// if l.X1 > l.X2 {
			// 	//xbig = l.X1
			// 	xsmall = l.X2
			// }

			for y := ysmall; y <= ybig; y++ {
				if g.Positions == nil {
					g.Positions = map[int]map[int]int{}
				}
				if g.Positions[y] == nil {
					g.Positions[y] = map[int]int{}
				}
				g.Positions[y][startX] += 1
				log.Printf("adding y[%d] x[%d]", y, startX)
				if startX > y {
					startX--
				} else {
					startX++
				}
			}

		}
	}
}

func (g *Grid) GetDangerousPositions() int {
	count := 0
	for _, r := range g.Positions {
		for _, c := range r {
			if c >= 2 {
				count++
			}
		}

	}
	return count
}

func getXY(s string) (int, int) {
	s = strings.TrimSpace(s)

	split := strings.Split(s, ",")

	x, _ := strconv.Atoi(string(split[0]))
	y, _ := strconv.Atoi(string(split[1]))

	return x, y

}
