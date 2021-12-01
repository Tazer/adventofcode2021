package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	//var version = flag.Int("version", 1, "first or second part of the assignment")

	flag.Parse()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []int{}

	for scanner.Scan() {
		i1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, i1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res := countIncrements(inputs)

	log.Printf("Result: %d", res)

}

func countIncrements(inputs []int) int {
	incCount := 0
	var prevValue int

	for _, i := range inputs {

		if prevValue == 0 {

			prevValue = i
			continue
		}
		if i > prevValue {
			incCount++
		}
		prevValue = i
	}
	return incCount
}
