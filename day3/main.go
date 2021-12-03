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

	log.Print(lines)

	res := GetGammaAndEpsilonRate(lines)

	log.Printf("Result %d", res.Decimal())

}

type GammaAndEpsilon struct {
	Gamma   string
	Epsilon string
}

func (g *GammaAndEpsilon) Add(i int) {
	if i == 1 {
		g.Gamma += "1"
		g.Epsilon += "0"
	} else {
		g.Gamma += "0"
		g.Epsilon += "1"
	}
}

func (g *GammaAndEpsilon) Decimal() int64 {
	log.Printf("gamma %s elips %s", g.Gamma, g.Epsilon)
	gamma, _ := strconv.ParseInt(g.Gamma, 2, 64)
	epsilon, _ := strconv.ParseInt(g.Epsilon, 2, 64)
	log.Printf("gamma %d elips %d", gamma, epsilon)
	return gamma * epsilon
}

func GetGammaAndEpsilonRate(input map[int][]int) GammaAndEpsilon {
	res := GammaAndEpsilon{}
	i := 0
	for {
		if l, ok := input[i]; ok {
			ones := 0
			zero := 0
			for _, v := range l {
				if v == 1 {
					ones++
				} else {
					zero++
				}
			}
			if ones > zero {
				res.Add(1)
			} else {
				res.Add(0)
			}
		} else {
			break
		}
		i++
	}
	return res
}
