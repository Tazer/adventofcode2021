package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	BingoGame struct {
		numbersDrawn []int
		bingoBoards  []BingoBoard
	}

	BingoBoard struct {
		Numbers map[int]map[int]BingoNumber
		HaveWon bool
	}

	BingoNumber struct {
		Drawn bool
		X, Y  int
		Value int
	}
)

func NewBingoGame(numbersDrawn []int, bingoBoards []BingoBoard) *BingoGame {
	return &BingoGame{
		numbersDrawn: numbersDrawn,
		bingoBoards:  bingoBoards,
	}
}

func (bb *BingoBoard) Mark(number int) {
	if bb.HaveWon {
		return
	}
	for x, row := range bb.Numbers {
		for y, bn := range row {
			if bn.Value == number {
				v := bb.Numbers[x][y]
				v.Drawn = true
				bb.Numbers[x][y] = v
				break
			}
		}
	}
}

func (bb *BingoBoard) Won() bool {

	if bb.HaveWon {
		return false
	}

	correct := 0
	winning := false
	for y := 0; y < 5; y++ {
		correct = 0
		for x := 0; x < 5; x++ {
			if bb.Numbers[x][y].Drawn {
				correct++
			}
			if correct == 5 {
				winning = true
			}
		}
	}

	if winning {
		bb.HaveWon = true
		return true
	}
	for x := 0; x < 5; x++ {
		correct = 0
		for y := 0; y < 5; y++ {

			if bb.Numbers[x][y].Drawn {
				correct++
			}
			if correct == 5 {
				winning = true
			}
		}
	}
	if winning {
		bb.HaveWon = true
		return true
	}
	return false
}

func (bb *BingoBoard) Score(winningNumber int) int {
	score := 0

	for _, row := range bb.Numbers {
		for _, bn := range row {
			if !bn.Drawn {
				score += bn.Value
			}
		}
	}

	return score * winningNumber
}

func (b *BingoGame) Play() int {
	for _, d := range b.numbersDrawn {
		for _, bb := range b.bingoBoards {
			bb.Mark(d)
			if bb.Won() {
				return bb.Score(d)
			}
		}
	}
	return 0
}

type WinningBoard struct {
	Number int
	Board  BingoBoard
}

func (w *WinningBoard) Score() int {
	return w.Board.Score(w.Number)
}

func (b *BingoGame) PlayLastWinner() int {

	boardsWinning := []WinningBoard{}

	for _, d := range b.numbersDrawn {
		for i, bb := range b.bingoBoards {
			b.bingoBoards[i].Mark(d)
			if b.bingoBoards[i].Won() {
				boardsWinning = append(boardsWinning, WinningBoard{
					Number: d,
					Board:  bb,
				})
			}
		}
	}

	return boardsWinning[len(boardsWinning)-1].Score()
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	b := parseBingoGame(scanner)

	s := b.Play()

	log.Printf("Played bingo %d", s)

	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	b = parseBingoGame(scanner)

	s = b.PlayLastWinner()

	log.Printf("Played bingo last winner %d", s)
}

func parseBingoGame(scanner *bufio.Scanner) BingoGame {
	drawnNumbers := []int{}
	bingoBoards := []BingoBoard{}

	addToDrawNumbers := true

	var curBoard BingoBoard
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			addToDrawNumbers = false
			if len(curBoard.Numbers) > 0 {
				bingoBoards = append(bingoBoards, curBoard)
			}
			curBoard = BingoBoard{}
			y = 0
			continue
		}

		if addToDrawNumbers {
			sLine := strings.Split(line, ",")

			for _, s := range sLine {
				i, _ := strconv.Atoi(s)

				drawnNumbers = append(drawnNumbers, i)
			}
		} else {
			sLine := strings.Split(line, " ")
			x := 0
			for _, s := range sLine {
				if s == "" {
					continue
				}
				i, err := strconv.Atoi(s)

				if err != nil {
					log.Fatal(err)
				}

				bingoNumber := BingoNumber{
					Drawn: false,
					X:     x,
					Y:     y,
					Value: i,
				}

				if curBoard.Numbers == nil {
					curBoard.Numbers = map[int]map[int]BingoNumber{}
				}

				if curBoard.Numbers[x] == nil {
					curBoard.Numbers[x] = map[int]BingoNumber{}
				}

				curBoard.Numbers[x][y] = bingoNumber
				x++
			}
			y++
		}

	}
	bingoBoards = append(bingoBoards, curBoard)
	return *NewBingoGame(drawnNumbers, bingoBoards)
}
