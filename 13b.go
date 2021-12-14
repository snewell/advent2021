package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

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

func makeOutput(coordinates map[point]struct{}) []string {
	output := []string{}
	for p := range coordinates {
		if p.y >= len(output) {
			neededLines := p.y - len(output)
			for i := 0; i < (neededLines + 1); i++ {
				output = append(output, "")
			}
		}

		if p.x >= len(output[p.y]) {
			neededPadding := p.x - len(output[p.y])
			output[p.y] += strings.Repeat(" ", neededPadding+1)
		}
		temp := []rune(output[p.y])
		temp[p.x] = '#'
		output[p.y] = string(temp)
	}
	return output
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		coordinates, instructions := readData(input)
		for _, inst := range instructions {
			coordinates = apply(coordinates, inst)
		}
		output := makeOutput(coordinates)
		for index := range output {
			fmt.Printf("%v\n", output[index])
		}
		return nil
	})
}
