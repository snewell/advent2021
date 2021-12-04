package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/snewell/advent2021/internal/aoc"
)

type Board [5][5]int

func readSelections(input io.Reader) []int {
	var raw string
	_, err := fmt.Fscanln(input, &raw)
	if err != nil {
		panic(err)
	}
	stringSelections := strings.Split(raw, ",")
	ret := make([]int, len(stringSelections))
	for index, value := range stringSelections {
		ret[index], err = strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
	}
	return ret
}

func readBoard(input io.Reader) (Board, error) {
	var ret Board
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			_, err := fmt.Fscan(input, &ret[i][j])
			if err != nil {
				return ret, err
			}
		}
	}
	return ret, nil
}

func checkRow(board *Board, row int) bool {
	for _, cell := range board[row] {
		if cell != -1 {
			return false
		}
	}
	return true
}

func checkColumn(board *Board, col int) bool {
	for row, _ := range board {
		if board[row][col] != -1 {
			return false
		}
	}
	return true
}

func checkBoard(board *Board, selection int) bool {
	for row, _ := range board {
		for col, cell := range board[row] {
			if cell == selection {
				board[row][col] = -1
				if checkRow(board, row) {
					return true
				}
				if checkColumn(board, col) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		selections := readSelections(input)

		boards := []*Board{}
		for {
			board, err := readBoard(input)
			if err != nil {
				break
			}
			boards = append(boards, &board)
		}

		for _, selection := range selections {
			for _, board := range boards {
				done := checkBoard(board, selection)
				if done {
					sum := 0
					for row := 0; row < 5; row++ {
						for col := 0; col < 5; col++ {
							if board[row][col] != -1 {
								sum += board[row][col]
							}
						}
					}
					return sum * selection
				}
			}
		}

		panic("No winner found")
	})
}
