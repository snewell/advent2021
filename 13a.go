package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"

	"github.com/snewell/advent2021/internal/aoc"
)

type point struct {
	x int
	y int
}

type instruction struct {
	horizontal bool
	axis       int
}

func readData(input io.Reader) (map[point]struct{}, []instruction) {
	coordinatePattern, _ := regexp.Compile(`^([0-9]+),([0-9]+)$`)
	coordinates := map[point]struct{}{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		subs := coordinatePattern.FindStringSubmatch(scanner.Text())
		x, _ := strconv.Atoi(subs[1])
		y, _ := strconv.Atoi(subs[2])
		coordinates[point{x: x, y: y}] = struct{}{}
	}

	instructionPattern, _ := regexp.Compile(`^fold along ([xy])=([0-9]+)$`)
	instructions := []instruction{}
	for scanner.Scan() {
		subs := instructionPattern.FindStringSubmatch(scanner.Text())
		axis, _ := strconv.Atoi(subs[2])
		instructions = append(instructions, instruction{
			horizontal: subs[1] == "y",
			axis:       axis,
		})
	}

	return coordinates, instructions
}

func apply(coordinates map[point]struct{}, instruction instruction) map[point]struct{} {
	coordinateMapper := func() func(point) point {
		axis := instruction.axis
		if instruction.horizontal {
			return func(p point) point {
				if p.y < instruction.axis {
					return p
				}
				return point{x: p.x, y: axis - (p.y - axis)}
			}
		}
		return func(p point) point {
			if p.x < axis {
				return p
			}
			return point{x: axis - (p.x - axis), y: p.y}
		}
	}()

	next := map[point]struct{}{}
	for p := range coordinates {
		next[coordinateMapper(p)] = struct{}{}
	}
	return next
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		coordinates, instructions := readData(input)
		next := apply(coordinates, instructions[0])
		return len(next)
	})
}
