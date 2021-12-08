package main

import (
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		positions := aoc.ReadCrabPositions(input)

		costFunction := func(distance int) int {
			return (distance*distance + distance) / 2
		}
		return aoc.CalculateMinFuel(positions, costFunction)
	})
}
