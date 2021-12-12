package main

import (
	"bufio"
	"io"
	"sort"

	"github.com/snewell/advent2021/internal/aoc"
)

func scoreCorrections(unclosed []rune) int {
	score := 0
	for index := range unclosed {
		score *= 5
		switch unclosed[len(unclosed)-index-1] {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}
	return score
}

func filterInput(input io.Reader, filter func(string) []rune) []int {
	scores := []int{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		unclosed := filter(scanner.Text())
		if len(unclosed) != 0 {
			scores = append(scores, scoreCorrections(unclosed))
		}
	}
	return scores
}

func isCorrupt(expected rune, stack []rune) bool {
	if len(stack) < 0 {
		return true
	}
	if stack[len(stack)-1] != expected {
		return true
	}
	return false
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		isCorrupt := func(line string) []rune {
			stack := []rune{}
			for _, ch := range line {
				switch ch {
				case ')':
					if isCorrupt('(', stack) {
						return nil
					}
					stack = stack[:len(stack)-1]
				case ']':
					if isCorrupt('[', stack) {
						return nil
					}
					stack = stack[:len(stack)-1]
				case '}':
					if isCorrupt('{', stack) {
						return nil
					}
					stack = stack[:len(stack)-1]
				case '>':
					if isCorrupt('<', stack) {
						return nil
					}
					stack = stack[:len(stack)-1]
				default:
					stack = append(stack, ch)
				}
			}
			return stack
		}

		scores := filterInput(input, isCorrupt)
		sort.Ints(scores)
		return scores[len(scores)/2]
	})
}
