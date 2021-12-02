package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

type Movement struct {
	direction string
	steps     int
}

type Position struct {
	horizontal int
	depth      int
	aim        int
}

func (m *Position) Result() int {
	return m.horizontal * m.depth
}

func main() {
	//var version = flag.Int("version", 1, "first or second part of the assignment")

	flag.Parse()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []Movement{}

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		iSplit, _ := strconv.Atoi(split[1])

		inputs = append(inputs, Movement{split[0], iSplit})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res := CalculateMovements(inputs, false)

	log.Printf("Result: %d", res.Result())

	res2 := CalculateMovements(inputs, true)

	log.Printf("Result2: %d", res2.Result())

}

func CalculateMovements(i []Movement, useAim bool) Position {

	pos := Position{0, 0, 0}

	for _, m := range i {

		switch m.direction {
		case "forward":
			pos.horizontal += m.steps
			if useAim {
				pos.depth += (pos.aim * m.steps)
			}
		case "down":

			if useAim {
				pos.aim += m.steps
			} else {
				pos.depth += m.steps
			}
		case "up":
			if useAim {
				pos.aim -= m.steps
			} else {
				pos.depth -= m.steps
			}
		}
	}
	return pos
}
