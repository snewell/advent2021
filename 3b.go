package main

import (
	"fmt"
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func stripValues(stripCounts []int, bitIndex int, current string) {
	for index, value := range current[bitIndex+1:] {
		if value == '1' {
			stripCounts[bitIndex+index+1]++
		}
	}
}

func findValue(raw []string, oneCounts []int, desiredBit func(int, int) byte) int {
	candidates := make([]string, len(raw))
	candidateCounts := make([]int, len(oneCounts))
	copy(candidates, raw)
	copy(candidateCounts, oneCounts)

	bitIndex := 0
	for len(candidates) > 1 {
		stripCounts := make([]int, len(candidateCounts))
		oneBits := candidateCounts[bitIndex]
		zeroBits := len(candidates) - candidateCounts[bitIndex]
		filtered := []string{}

		desired := desiredBit(oneBits, zeroBits)
		for _, current := range candidates {
			if current[bitIndex] == desired {
				filtered = append(filtered, current)
			} else {
				stripValues(stripCounts, bitIndex, current)
			}
		}

		for index, _ := range candidateCounts {
			candidateCounts[index] -= stripCounts[index]
		}

		candidates = filtered
		bitIndex++
	}

	ret := 0
	for _, value := range candidates[0] {
		ret <<= 1
		if value == '1' {
			ret |= 1
		}
	}
	return ret
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		var current string
		fmt.Fscan(input, &current)
		oneCounts := make([]int, len(current))
		rawData := []string{}
		for {
			rawData = append(rawData, current)
			for index, value := range current {
				if value == '1' {
					oneCounts[index]++
				}
			}

			_, err := fmt.Fscan(input, &current)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}

		oxygenValue := findValue(rawData, oneCounts, func(ones int, zeros int) byte {
			if ones >= zeros {
				return '1'
			}
			return '0'
		})
		co2Value := findValue(rawData, oneCounts, func(ones int, zeros int) byte {
			if ones >= zeros {
				return '0'
			}
			return '1'
		})
		return oxygenValue * co2Value
	})
}
