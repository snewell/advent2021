package aoc

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type CrabPositionData struct {
	Counts      map[int]int
	MinPosition int
	MaxPosition int
}

func ReadCrabPositions(input io.Reader) CrabPositionData {
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

	return CrabPositionData{
		Counts:      counts,
		MinPosition: minPosition,
		MaxPosition: maxPosition,
	}
}

func CalculateMinFuel(positions CrabPositionData, costFunction func(int) int) int {
	minFuel := math.MaxInt32
	for candidate := positions.MinPosition; candidate <= positions.MaxPosition; candidate++ {
		totalFuel := 0
		for position, count := range positions.Counts {
			distance := candidate - position
			if distance < 0 {
				distance *= -1
			}
			totalFuel += costFunction(distance) * count
		}

		if totalFuel < minFuel {
			minFuel = totalFuel
		}
	}
	return minFuel
}
