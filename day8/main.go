package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
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

	log.Print(w)
}

func parseLine(line string) *Line {
	s := strings.Split(line, "|")
	st1 := strings.TrimSpace(s[0])
	ps := strings.Split(st1, " ")

	st2 := strings.TrimSpace(s[1])
	os := strings.Split(st2, " ")
	return &Line{
		Pattern: ps,
		Output:  os,
	}

}

func NewWorld(lines []*Line) *World {
	return &World{
		Lines: lines,
	}
}

type World struct {
	Lines []*Line
}

type Line struct {
	Pattern []string
	Output  []string
}
