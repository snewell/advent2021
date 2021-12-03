package main

import (
	"fmt"
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		var current string
		fmt.Fscan(input, &current)
		oneCounts := make([]int, len(current))
		total := 0
		for {
			total++
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

		oneLimit := total / 2

		gammaRate := 0
		epsilonRate := 0
		for _, value := range oneCounts {
			gammaRate <<= 1
			epsilonRate <<= 1
			if value > oneLimit {
				gammaRate |= 1
			} else {
				epsilonRate |= 1
			}
		}

		return gammaRate * epsilonRate
	})
}
