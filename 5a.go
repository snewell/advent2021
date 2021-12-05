package main

import (
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func isHorizontal(line aoc.Line) bool {
	return line.First.Y == line.Second.Y
}

func filter(lines []aoc.Line) ([]aoc.Line, []aoc.Line) {
	hor := []aoc.Line{}
	vert := []aoc.Line{}

	for _, line := range lines {
		if isHorizontal(line) {
			if line.First.X > line.Second.X {
				line = aoc.Line{
					First:  line.Second,
					Second: line.First,
				}
			}
			hor = append(hor, line)
		} else {
			if line.First.Y > line.Second.Y {
				line = aoc.Line{
					First:  line.Second,
					Second: line.First,
				}
			}
			vert = append(vert, line)
		}
	}

	return hor, vert
}

func findIntersection(hor aoc.Line, vert aoc.Line) (aoc.Point, bool) {
	if (hor.First.X <= vert.First.X) && (vert.First.X <= hor.Second.X) {
		// X overlap
		if (vert.First.Y <= hor.First.Y) && (hor.First.Y <= vert.Second.Y) {
			// Y overlap
			return aoc.Point{
				X: vert.First.X,
				Y: hor.First.Y,
			}, true
		}
	}
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

		// look for intersections
		hor, vert := filter(lines)
		counts := map[aoc.Point]int{}
		for _, horLine := range hor {
			for _, vertLine := range vert {
				point, intersects := findIntersection(horLine, vertLine)
				if intersects {
					counts[point] += 2
				}
			}
		}

		// check for overlapping horizontal lines
		for index, first := range hor[:len(hor)-1] {
			for _, second := range hor[index+1:] {
				overlapping := aoc.FindHorizontalOverlap(first, second)
				for _, point := range overlapping {
					counts[point] += 2
				}
			}
		}

		// check for overlapping vertical lines
		for index, first := range vert[:len(vert)-1] {
			for _, second := range vert[index+1:] {
				overlapping := aoc.FindVerticalOverlap(first, second)
				for _, point := range overlapping {
					counts[point] += 2
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
