package main

import (
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		positions := aoc.ReadCrabPositions(input)

		distanceCosts := map[int]int{}
		maxDistance := positions.MaxPosition - positions.MinPosition
		currentCost := 0
		for i := 0; i < maxDistance; i++ {
			currentCost += i
			distanceCosts[i] = currentCost
		}

		costFunction := func(distance int) int {
			return distanceCosts[distance]
		}
		return aoc.CalculateMinFuel(positions, costFunction)
	})
}
