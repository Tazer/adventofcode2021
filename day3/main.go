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

	log.Print(lines)

	res := GetGammaAndEpsilonRate(lines)

	log.Printf("Result %d", res.Decimal())

	resLife := GetLifeSupport(sLines)

	log.Printf("LIFE SUPPORT %d", resLife.Decimal())

}

type GammaAndEpsilon struct {
	Gamma   string
	Epsilon string
}

type LifeSupportRating struct {
	OxygenGeneratorRating string
	CO2ScrubbingRating    string
}

func (l *LifeSupportRating) Decimal() int64 {
	oxygen, _ := strconv.ParseInt(l.OxygenGeneratorRating, 2, 64)
	co2, _ := strconv.ParseInt(l.CO2ScrubbingRating, 2, 64)
	log.Printf("oxygen %d co2 %d", oxygen, co2)
	return oxygen * co2
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

func GetLifeSupport(input []string) LifeSupportRating {
	res := LifeSupportRating{}
	//Oxygen
	resList := getLastOne(input, true, 0)
	log.Printf("Oxygen %v", resList)
	res.OxygenGeneratorRating = resList[0]

	//Co2
	resList = getLastOne(input, false, 0)
	log.Printf("co2 %v", resList)
	res.CO2ScrubbingRating = resList[0]

	return res
}

func getLastOne(input []string, most bool, index int) []string {

	if len(input) == 1 {
		return input
	}

	ones := 0
	zero := 0

	for _, v := range input {
		str := string(v[index])
		if str == "1" {
			ones++
		} else {
			zero++
		}
	}

	allowOnes := false

	if most {
		if ones > zero || ones == zero {
			allowOnes = true
		} else {
			allowOnes = false
		}
	} else {
		if ones < zero {
			allowOnes = true
		} else {
			allowOnes = false
		}
	}

	allowed := []string{}

	for _, v := range input {
		str := string(v[index])
		if str == "1" && allowOnes {
			allowed = append(allowed, v)
		} else if str == "0" && !allowOnes {
			allowed = append(allowed, v)
		}
	}

	return getLastOne(allowed, most, index+1)
}
