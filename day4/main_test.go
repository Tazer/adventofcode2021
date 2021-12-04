package main

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBingo(t *testing.T) {
	file, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	game := parseBingoGame(scanner)

	log.Printf("%+v", game)

	assert.Equal(t, 7, game.numbersDrawn[0])
	assert.Equal(t, 4, game.numbersDrawn[1])
	assert.Equal(t, 1, game.numbersDrawn[len(game.numbersDrawn)-1])

	assert.Equal(t, 22, game.bingoBoards[0].Numbers[0][0].Value)
	assert.Equal(t, 3, game.bingoBoards[1].Numbers[0][0].Value)
	assert.Equal(t, 14, game.bingoBoards[2].Numbers[0][0].Value)
	assert.Equal(t, 4, game.bingoBoards[2].Numbers[4][0].Value)
	assert.Equal(t, 4, game.bingoBoards[2].Numbers[4][0].X)
	assert.Equal(t, 0, game.bingoBoards[2].Numbers[4][0].Y)
}

func TestPlayBingo(t *testing.T) {
	file, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	game := parseBingoGame(scanner)

	s := game.Play()

	assert.Equal(t, 4512, s)
}
