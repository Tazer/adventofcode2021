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

	w := NewWorld(initState, 80, false)

	res := w.Simulate()

	log.Printf("Result: %d", res)

	w2 := NewWorld(initState, 256, true)

	res2 := w2.Simulate()

	log.Printf("Result: %d", res2)
}

func NewWorld(initState []int, days int, useFast bool) *World {

	fishes := []*Fish{}
	fm := make(map[int]int, 8)
	for _, a := range initState {
		f := NewFish(a, false)
		fishes = append(fishes, f)
		fm[a] += 1
	}

	return &World{
		Days:        days,
		fishes:      fishes,
		useFastImpl: useFast,
		fm:          fm,
	}
}

type World struct {
	fishes      []*Fish
	Day         int
	Days        int
	useFastImpl bool
	fm          map[int]int
}

func (w *World) TickFastImpl() {

	new := make(map[int]int, 8)

	for k, v := range w.fm {
		if k == 0 {
			new[6] += v
			new[8] += v
			continue
		}
		new[k-1] += v
	}
	w.fm = new
}

func (w *World) Tick() {
	newFishes := []*Fish{}
	for _, f := range w.fishes {
		nf := f.Tick()

		if nf != nil {
			newFishes = append(newFishes, nf)
		}
	}
	w.fishes = append(w.fishes, newFishes...)
}

func (w *World) Simulate() int {

	for w.Day = 1; w.Day <= w.Days; w.Day++ {
		if w.useFastImpl {
			w.TickFastImpl()
		} else {
			w.Tick()
		}
	}
	if w.useFastImpl {
		val := 0
		for _, v := range w.fm {
			val += v
		}
		return val
	}
	return len(w.fishes)
}

func NewFish(age int, newBorn bool) *Fish {
	return &Fish{
		Age:     age,
		NewBorn: newBorn,
	}
}

type Fish struct {
	Age     int
	NewBorn bool
}

func (f *Fish) Tick() *Fish {

	if f.Age == 0 {
		f.Age = 6
		return NewFish(8, true)
	}

	if f.Age > 0 {
		f.Age -= 1
	}

	return nil
}
