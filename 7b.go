package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/snewell/advent2021/internal/aoc"
)

type positionData struct {
	counts      map[int]int
	minPosition int
	maxPosition int
}

func readPositions(input io.Reader) positionData {
	counts := map[int]int{}
	var data string
	fmt.Fscan(input, &data)
	values := strings.Split(data, ",")

	minPosition := math.MaxInt32
	maxPosition := math.MinInt32
	for _, value := range values {
		realValue, _ := strconv.Atoi(value)
		if realValue < minPosition {
			minPosition = realValue
		}
		if maxPosition < realValue {
			maxPosition = realValue
		}
		counts[realValue]++
	}

	return positionData{
		counts:      counts,
		minPosition: minPosition,
		maxPosition: maxPosition,
	}
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		positions := readPositions(input)

		minFuel := math.MaxInt32
		for candidate := positions.minPosition; candidate <= positions.maxPosition; candidate++ {
			totalFuel := 0
			for position, count := range positions.counts {
				distance := candidate - position
				if distance < 0 {
					distance *= -1
				}
				cost := 0
				for ; distance > 0; distance-- {
					cost += distance
				}
				totalFuel += cost * count
			}

			if totalFuel < minFuel {
				minFuel = totalFuel
			}
		}
		return minFuel
	})
}
