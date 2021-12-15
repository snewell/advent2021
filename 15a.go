package main

import (
	"bufio"
	"io"
	"math"
	"sort"

	"github.com/snewell/advent2021/internal/aoc"
)

func loadMap(input io.Reader) [][]int {
	ret := [][]int{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for index := range line {
			row[index] = int(line[index] - '0')
		}
		ret = append(ret, row)

	}
	return ret
}

type point struct {
	x int
	y int
}

func dijkstra(riskMap [][]int) int {
	minRisk := make([][]int, len(riskMap))
	for row := range riskMap {
		minRisk[row] = make([]int, len(riskMap[row]))
		for col := range riskMap[row] {
			minRisk[row][col] = math.MaxInt32
		}
	}

	toVisit := []point{point{x: 0, y: 0}}
	minRisk[0][0] = 0
	goalRow := len(riskMap) - 1
	goalCol := len(riskMap[goalRow]) - 1
	for minRisk[goalRow][goalCol] == math.MaxInt32 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		localCost := minRisk[current.y][current.x]
		if current.y != 0 {
			targetRisk := riskMap[current.y-1][current.x] + localCost
			if targetRisk < minRisk[current.y-1][current.x] {
				minRisk[current.y-1][current.x] = targetRisk
				toVisit = append(toVisit, point{x: current.x, y: current.y - 1})
			}
		}
		if current.y != goalCol {
			targetRisk := riskMap[current.y+1][current.x] + localCost
			if targetRisk < minRisk[current.y+1][current.x] {
				minRisk[current.y+1][current.x] = targetRisk
				toVisit = append(toVisit, point{x: current.x, y: current.y + 1})
			}
		}

		if current.x != 0 {
			targetRisk := riskMap[current.y][current.x-1] + localCost
			if targetRisk < minRisk[current.y][current.x-1] {
				minRisk[current.y][current.x-1] = targetRisk
				toVisit = append(toVisit, point{x: current.x - 1, y: current.y})
			}
		}
		if current.x != goalRow {
			targetRisk := riskMap[current.y][current.x+1] + localCost
			if targetRisk < minRisk[current.y][current.x+1] {
				minRisk[current.y][current.x+1] = targetRisk
				toVisit = append(toVisit, point{x: current.x + 1, y: current.y})
			}
		}

		sort.Slice(toVisit, func(lhs int, rhs int) bool {
			toVisitLhs := toVisit[lhs]
			toVisitRhs := toVisit[rhs]
			return minRisk[toVisitLhs.y][toVisitLhs.x] < minRisk[toVisitRhs.y][toVisitRhs.x]
		})
	}

	return minRisk[goalRow][goalCol]
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		riskMap := loadMap(input)
		return dijkstra(riskMap)
	})
}
