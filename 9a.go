package main

import (
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		total := 0
		heightMap := aoc.LoadBasins(input)
		aoc.FindLowBasinPoints(heightMap, func(row int, col int) {
			total += heightMap[row][col] + 1
		})
		return total
	})
}
