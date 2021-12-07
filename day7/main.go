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

	initState := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		sLine := strings.Split(line, ",")

		for _, sl := range sLine {
			i, _ := strconv.Atoi(sl)
			initState = append(initState, i)

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	w := NewWorld(initState)

	res := w.Simulate(false)

	log.Printf("Best %d", res)

	res = w.Simulate(true)

	log.Printf("Best V2 %d", res)

}

func NewWorld(initState []int) *World {

	w := &World{
		Crabs: []Crab{},
	}

	for _, i := range initState {
		c := Crab{
			Position: i,
		}
		w.AddCrab(c)
	}
	return w
}

func (w *World) Simulate(moveV2 bool) int {

	bestFuel := 0

	for i := w.MinPosition; i <= w.MaxPosition; i++ {
		usedFuel := 0
		for _, c := range w.Crabs {
			if moveV2 {
				usedFuel += c.MoveV2(i)
			} else {
				usedFuel += c.Move(i)
			}
		}
		if usedFuel < bestFuel || bestFuel == 0 {
			bestFuel = usedFuel
		}
	}

	return bestFuel
}

func (w *World) AddCrab(c Crab) {
	if c.Position > w.MaxPosition {
		w.MaxPosition = c.Position
	}

	if c.Position < w.MinPosition || w.MinPosition == 0 {
		w.MinPosition = c.Position
	}

	w.Crabs = append(w.Crabs, c)
}

type World struct {
	MaxPosition int
	MinPosition int
	Crabs       []Crab
}

type Crab struct {
	Position int
	// FuelUsed int
}

func (c *Crab) Move(i int) int {

	if c.Position > i {
		return c.Position - i
	}
	return i - c.Position

}

func (c *Crab) MoveV2(m int) int {

	fuelUsed := 0

	small := m
	big := c.Position

	if small > big {
		small = c.Position
		big = m
	}

	cost := 0

	for i := small; i < big; i++ {
		fuelUsed += 1 + cost
		cost++
	}

	return fuelUsed
}
