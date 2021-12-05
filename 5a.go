package main

import (
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func isHorizontal(line aoc.Line) bool {
	return line.First.X == line.Second.X
}

func filter(lines []aoc.Line) ([]aoc.Line, []aoc.Line) {
	hor := []aoc.Line{}
	vert := []aoc.Line{}

	for _, line := range lines {
		if isHorizontal(line) {
			// TODO: order points in line
			hor = append(hor, line)
		} else {
			// TODO: order points in line
			vert = append(vert, line)
		}
	}

	return hor, vert
}

func findIntersection(hor aoc.Line, vert aoc.Line) (aoc.Point, bool) {
	// TODO: implement
	return aoc.Point{}, false
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		lines, err := aoc.ReadLines(input, func(line aoc.Line) bool {
			if line.First.X == line.Second.X {
				return true
			}
			if line.First.Y == line.Second.Y {
				return true
			}
			return false
		})
		if err != nil {
			panic(err)
		}

		hor, vert := filter(lines)
		counts := map[aoc.Point]int{}
		for _, horLine := range hor {
			for _, vertLine := range vert {
				point, intersects := findIntersection(horLine, vertLine)
				if intersects {
					counts[point]++
				}
			}
		}

		dangerousSpots := 0
		for _, count := range counts {
			if count >= 2 {
				dangerousSpots++
			}
		}

		return dangerousSpots
	})
}
