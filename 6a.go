package main

import (
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		counts := aoc.ReadFishPeriods(input)
		finalCounts := aoc.PassDays(counts, 80)

		total := 0
		for _, count := range finalCounts {
			total += count
		}
		return total
	})
}
