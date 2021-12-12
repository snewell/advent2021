package main

import (
	"bufio"
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func handleInput(input io.Reader, getScore func(string) int) int {
	score := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		score += getScore(scanner.Text())
	}
	return score
}

func matchChar(expected rune, stack []rune, score int) int {
	if len(stack) < 0 {
		return score
	}
	if stack[len(stack)-1] != expected {
		return score
	}
	return 0
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		scoreLine := func(line string) int {
			stack := []rune{}
			for _, ch := range line {
				switch ch {
				case ')':
					nextScore := matchChar('(', stack, 3)
					if nextScore != 0 {
						return nextScore
					}
					stack = stack[:len(stack)-1]
				case ']':
					nextScore := matchChar('[', stack, 57)
					if nextScore != 0 {
						return nextScore
					}
					stack = stack[:len(stack)-1]
				case '}':
					nextScore := matchChar('{', stack, 1197)
					if nextScore != 0 {
						return nextScore
					}
					stack = stack[:len(stack)-1]
				case '>':
					nextScore := matchChar('<', stack, 25137)
					if nextScore != 0 {
						return nextScore
					}
					stack = stack[:len(stack)-1]
				default:
					stack = append(stack, ch)
				}
			}
			return 0
		}
		return handleInput(input, scoreLine)
	})
}
